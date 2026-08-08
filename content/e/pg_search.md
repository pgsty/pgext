---
title: "pg_search"
linkTitle: "pg_search"
description: "Full text search for PostgreSQL using BM25"
weight: 2100
categories: ["FTS"]
width: full
---

[**pg_search**](https://github.com/paradedb/paradedb/tree/main/pg_search) : Full text search for PostgreSQL using BM25


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/main/pg_search" >}} | {{< ext "pg_search" >}} | `0.25.1` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `paradedb` |
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "pgfaceting" >}} {{< ext "roaringbitmap" >}} {{< ext "rum" >}} |

> [!Note] Requires shared_preload_libraries=pg_search and pgvector; bm25 access method conflicts with pg_textsearch and vchord_bm25; PIGSTY uses pgrx 0.19.1 for upstream pgrx 0.19.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.25.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_search` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.25.1` | {{< bg "18" "pg_search_18" "green" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "red" >}} | `pg_search_$v` | `pgvector_$v`, `openblas` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.25.1` | {{< bg "18" "postgresql-18-pg-search" "green" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "red" >}} | `postgresql-$v-pg-search` | `postgresql-$v-pgvector`, `libopenblas0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.25.1" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_18` | `0.25.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.5 MiB | [pg_search_18-0.25.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_18-0.25.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_18` | `0.25.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 69.5 MiB | [pg_search_18-0.25.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_18-0.25.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_18` | `0.25.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 70.6 MiB | [pg_search_18-0.25.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_18-0.25.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_18` | `0.25.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.5 MiB | [pg_search_18-0.25.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_18-0.25.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_18` | `0.25.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.5 MiB | [pg_search_18-0.25.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_18-0.25.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_18` | `0.25.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 69.4 MiB | [pg_search_18-0.25.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_18-0.25.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-search` | `0.25.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.0 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 63.8 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.0 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 63.8 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.0 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 67.3 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.9 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.1 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.8 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-search` | `0.25.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 67.0 MiB | [postgresql-18-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.25.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.6 MiB | [pg_search_17-0.25.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.25.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_17` | `0.25.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 69.5 MiB | [pg_search_17-0.25.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.25.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_17` | `0.25.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 70.6 MiB | [pg_search_17-0.25.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.25.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_17` | `0.25.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.4 MiB | [pg_search_17-0.25.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.25.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_17` | `0.25.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.6 MiB | [pg_search_17-0.25.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_17-0.25.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_17` | `0.25.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 69.4 MiB | [pg_search_17-0.25.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_17-0.25.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.25.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.0 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 63.8 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.0 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 63.8 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.0 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 67.4 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.9 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.1 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.8 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-search` | `0.25.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 67.1 MiB | [postgresql-17-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.25.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.6 MiB | [pg_search_16-0.25.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.25.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_16` | `0.25.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 69.5 MiB | [pg_search_16-0.25.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.25.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_16` | `0.25.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 70.7 MiB | [pg_search_16-0.25.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.25.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_16` | `0.25.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.5 MiB | [pg_search_16-0.25.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.25.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_16` | `0.25.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.6 MiB | [pg_search_16-0.25.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_16-0.25.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_16` | `0.25.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 69.3 MiB | [pg_search_16-0.25.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_16-0.25.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.25.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.0 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 63.8 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.0 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 63.8 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.0 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 67.2 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.9 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.1 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.8 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-search` | `0.25.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 67.0 MiB | [postgresql-16-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.25.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.5 MiB | [pg_search_15-0.25.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.25.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_15` | `0.25.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 69.5 MiB | [pg_search_15-0.25.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.25.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_15` | `0.25.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 70.6 MiB | [pg_search_15-0.25.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.25.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_search_15` | `0.25.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.4 MiB | [pg_search_15-0.25.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.25.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_search_15` | `0.25.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.6 MiB | [pg_search_15-0.25.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_15-0.25.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_search_15` | `0.25.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 69.3 MiB | [pg_search_15-0.25.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_15-0.25.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.25.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.0 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 63.7 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.0 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 63.8 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.0 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 67.2 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.9 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.1 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.8 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-search` | `0.25.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 67.0 MiB | [postgresql-15-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.25.1-1PIGSTY~resolute_arm64.deb) |

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
{{< card link="https://github.com/paradedb/paradedb/tree/main/pg_search" title="Repository" icon="github" subtitle="github.com/paradedb/paradedb/tree/main/pg_search" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.25.1.tar.gz" >}}
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
CREATE EXTENSION pg_search CASCADE; -- requires vector
```

## Usage

Sources:

- [pg_search v0.25.1 README](https://github.com/paradedb/paradedb/blob/v0.25.1/pg_search/README.md)
- [pg_search v0.25.1 release](https://github.com/paradedb/paradedb/releases/tag/v0.25.1)
- [pg_search v0.25.1 changelog](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/changelog/0.25.1.mdx)
- [Create a ParadeDB index](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/indexing/create-index.mdx)
- [Full-text match operators](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/full-text/match.mdx)
- [BM25 scoring](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/sorting/score.mdx)
- [Highlighting and snippets](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/full-text/highlight.mdx)
- [Index vectors](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/indexing/indexing-vectors.mdx)
- [Query vectors](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/vector/querying.mdx)
- [Hybrid-search overview](https://github.com/paradedb/paradedb/blob/v0.25.1/docs/documentation/hybrid/overview.mdx)

`pg_search` adds ParadeDB's full-text, structured, vector, and hybrid search index to PostgreSQL. Version 0.25 uses the `paradedb` index access method; the older `bm25` access-method name remains a compatibility alias. The extension requires `vector`, supports PostgreSQL 15-18 upstream, and must be loaded through `shared_preload_libraries`.

### Install and Build an Index

```conf
shared_preload_libraries = 'pg_search'
```

Restart PostgreSQL, then create the extension and a table with a stable unique key:

```sql
CREATE EXTENSION pg_search CASCADE;

CREATE TABLE documents (
  id          bigint PRIMARY KEY,
  title       text,
  body        text,
  category    text,
  embedding   vector(768)
);

CREATE INDEX documents_search_idx ON documents
USING paradedb (
  id,
  title,
  body,
  category,
  embedding vector_cosine_ops
)
WITH (key_field = 'id');
```

The `key_field` must be the first indexed column and uniquely identify every row. A text key must be indexed without tokenization. A table can have only one ParadeDB index, so include every searchable field in that index.

### Full-Text Search

Use `|||` to match any token and `&&&` to require all tokens:

```sql
SELECT id, title, pdb.score(id) AS score
FROM documents
WHERE body ||| 'postgresql search'
ORDER BY score DESC, id;

SELECT id, pdb.snippet(body) AS excerpt
FROM documents
WHERE body &&& 'postgresql indexing';
```

`pdb.score(key_field)` exposes the relevance score for the current row. `pdb.snippet(indexed_text_column)` returns a highlighted excerpt. These helpers are meaningful only in a query driven by a ParadeDB search predicate.

### Vector Search

Vector indexing is beta in the 0.25 line and uses the `vector` type from pgvector. Choose the operator class when the index is created; changing the metric requires rebuilding the index.

```sql
SELECT id, title, embedding <=> $1::vector AS distance
FROM documents
WHERE id @@@ pdb.all()
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

Supported index operator classes are `vector_l2_ops`, `vector_ip_ops`, and `vector_cosine_ops`. The 0.25 vector index does not index `halfvec`, `sparsevec`, or `bit` columns.

### Hybrid Search

A single ParadeDB index can combine lexical predicates, structured filters, and vector ordering. For more elaborate fusion, use the documented RRF and weighted hybrid-search functions instead of adding scores from unrelated scales directly.

```sql
SELECT id, title, pdb.score(id) AS lexical_score
FROM documents
WHERE body ||| 'postgresql extension'
  AND category === 'database'
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

### Version 0.25.1 and Caveats

- Version 0.25 renamed the primary index access method from `bm25` to `paradedb`. Existing `USING bm25` definitions remain supported, but new examples should use `USING paradedb`.
- Version 0.25.1 supports deterministic vector tie breakers and pushes the vector arm of reciprocal-rank-fusion queries into the index. It also adds `paradedb.vector_clustering_threshold`, whose default is 500, and caps vector-index build parallelism at four workers.
- Version 0.25.1 removes `paradedb.vector_cluster_probe_epsilon` and changes the vector-index bounds gate. After upgrading a database from 0.25.0, `REINDEX` every ParadeDB index that contains a vector field; installing the new shared library and running `ALTER EXTENSION` alone is not sufficient for those indexes.
- `CREATE EXTENSION pg_search CASCADE` can install the required `vector` extension, but every server process still needs the preload configuration and restart first.
- Query plans, tokenization, and ranking can change when an index is rebuilt with different field options. Test relevance and vector recall with production-shaped data before rollout.
