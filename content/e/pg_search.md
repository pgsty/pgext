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
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< ext "pg_search" >}} | `0.22.6` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `paradedb` |
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} |

> [!Note] bm25 am conflict with pg_textsearch, PG 17+ does not require dynamic loading 


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.22.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_search` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.22.6` | {{< bg "18" "pg_search_18" "green" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "red" >}} | `pg_search_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.22.6` | {{< bg "18" "postgresql-18-pg-search" "green" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "red" >}} | `postgresql-$v-pg-search` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.10" "pg_search_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "pg_search_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.5" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.6" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_18` | `0.22.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 37.7 MiB | [pg_search_18-0.22.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_18-0.22.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_18` | `0.22.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 MiB | [pg_search_18-0.22.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_18-0.22.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_18` | `0.22.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.4 MiB | [pg_search_18-0.22.6-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_18-0.22.6-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_18` | `0.22.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.1 MiB | [pg_search_18-0.22.6-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_18-0.22.6-1PARADEDB.el9.aarch64.rpm) |
| `pg_search_18` | `0.22.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.5 MiB | [pg_search_18-0.22.6-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_18-0.22.6-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_18` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.0 MiB | [pg_search_18-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_18-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-18-pg-search` | `0.22.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.0 MiB | [postgresql-18-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_amd64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.4 MiB | [postgresql-18-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_arm64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 58.9 MiB | [postgresql-18-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_amd64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.4 MiB | [postgresql-18-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_arm64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.0 MiB | [postgresql-18-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_amd64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.4 MiB | [postgresql-18-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_arm64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 58.9 MiB | [postgresql-18-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_amd64.deb) |
| `postgresql-18-pg-search` | `0.22.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.5 MiB | [postgresql-18-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-18-pg-search/postgresql-18-pg-search_0.22.6_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.22.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.9 MiB | [pg_search_17-0.22.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.22.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_17` | `0.22.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 MiB | [pg_search_17-0.22.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.22.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_17` | `0.22.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.4 MiB | [pg_search_17-0.22.6-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.22.6-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_17` | `0.22.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.1 MiB | [pg_search_17-0.22.6-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.22.6-1PARADEDB.el9.aarch64.rpm) |
| `pg_search_17` | `0.22.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.4 MiB | [pg_search_17-0.22.6-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_17-0.22.6-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_17` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 MiB | [pg_search_17-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_17-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.22.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.0 MiB | [postgresql-17-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_amd64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.4 MiB | [postgresql-17-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_arm64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.0 MiB | [postgresql-17-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_amd64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.4 MiB | [postgresql-17-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_arm64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.0 MiB | [postgresql-17-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_amd64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.4 MiB | [postgresql-17-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_arm64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 59.0 MiB | [postgresql-17-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_amd64.deb) |
| `postgresql-17-pg-search` | `0.22.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.4 MiB | [postgresql-17-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-17-pg-search/postgresql-17-pg-search_0.22.6_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.22.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.9 MiB | [pg_search_16-0.22.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.22.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_16` | `0.22.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 MiB | [pg_search_16-0.22.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.22.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_16` | `0.22.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.4 MiB | [pg_search_16-0.22.6-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.22.6-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_16` | `0.22.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.0 MiB | [pg_search_16-0.22.6-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.22.6-1PARADEDB.el9.aarch64.rpm) |
| `pg_search_16` | `0.22.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.4 MiB | [pg_search_16-0.22.6-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_16-0.22.6-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_16` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 MiB | [pg_search_16-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_16-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.22.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.0 MiB | [postgresql-16-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_amd64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.4 MiB | [postgresql-16-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_arm64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.0 MiB | [postgresql-16-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_amd64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.4 MiB | [postgresql-16-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_arm64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.0 MiB | [postgresql-16-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_amd64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.4 MiB | [postgresql-16-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_arm64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 59.0 MiB | [postgresql-16-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_amd64.deb) |
| `postgresql-16-pg-search` | `0.22.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.4 MiB | [postgresql-16-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-16-pg-search/postgresql-16-pg-search_0.22.6_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.22.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 60.9 MiB | [pg_search_15-0.22.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.22.6-1PIGSTY.el8.x86_64.rpm) |
| `pg_search_15` | `0.22.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 MiB | [pg_search_15-0.22.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.22.6-1PIGSTY.el8.aarch64.rpm) |
| `pg_search_15` | `0.22.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.5 MiB | [pg_search_15-0.22.6-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.22.6-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_15` | `0.22.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.0 MiB | [pg_search_15-0.22.6-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.22.6-1PARADEDB.el9.aarch64.rpm) |
| `pg_search_15` | `0.22.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.4 MiB | [pg_search_15-0.22.6-1PARADEDB.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_15-0.22.6-1PARADEDB.el10.x86_64.rpm) |
| `pg_search_15` | `0.22.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 MiB | [pg_search_15-0.22.6-1PARADEDB.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_15-0.22.6-1PARADEDB.el10.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.22.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 58.9 MiB | [postgresql-15-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_amd64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 58.4 MiB | [postgresql-15-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_arm64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.0 MiB | [postgresql-15-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_amd64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 58.4 MiB | [postgresql-15-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_arm64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 58.9 MiB | [postgresql-15-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_amd64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.4 MiB | [postgresql-15-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_arm64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 58.9 MiB | [postgresql-15-pg-search_0.22.6_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_amd64.deb) |
| `postgresql-15-pg-search` | `0.22.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 58.4 MiB | [postgresql-15-pg-search_0.22.6_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresql-15-pg-search/postgresql-15-pg-search_0.22.6_arm64.deb) |

{{< /tab >}}
{{< tab >}}

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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.22.5.tar.gz" >}}
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


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_search;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_search;
> CREATE INDEX search_idx ON mock_items
> USING bm25 (id, description, category)
> WITH (key_field='id');
> SELECT * FROM mock_items WHERE description @@@ 'keyboard';
> ```
>
> Sources: [README](https://github.com/paradedb/paradedb/tree/dev/pg_search), [Quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart), [Install docs](https://docs.paradedb.com/documentation/getting-started/install)

`pg_search` is ParadeDB's full text search extension for PostgreSQL. It provides BM25-based indexing and querying on heap tables, is built on Tantivy, and the current upstream README states support starts at PostgreSQL 15.

## Setup

The upstream installation docs highlight one critical requirement: `pg_search` must be included in `shared_preload_libraries` so its background worker can process index writes.

```ini
shared_preload_libraries = 'pg_search'
```

After that:

```sql
CREATE EXTENSION pg_search;
ALTER SYSTEM SET paradedb.pg_search_telemetry TO 'off';
```

## Creating a BM25 Index

The quickstart demonstrates creating a BM25 index on a heap table, with a unique key field:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating, in_stock, created_at, metadata, weight_range)
WITH (key_field='id');
```

The existing docs emphasize that the `key_field` must be unique and that BM25 indexes are the core access method for search queries.

## Querying

The `@@@` operator performs search queries:

```sql
SELECT description, rating, category
FROM mock_items
WHERE description @@@ 'keyboard' AND rating > 2
ORDER BY rating
LIMIT 5;
```

ParadeDB also documents helper functions for relevance scoring and snippets:

```sql
SELECT description, paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;

SELECT description, paradedb.snippet(description), paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;
```

Phrase search is supported with quoted expressions:

```sql
SELECT description
FROM mock_items
WHERE description @@@ '"metal keyboard"';
```

## Text Configuration

The quickstart also shows that text fields can be wrapped with tokenizer configuration, for example English stemming:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, (description::pdb.simple('stemmer=english')), category)
WITH (key_field='id');
```

For deeper setup and operational guidance, the upstream project points to the ParadeDB documentation site as the primary docs surface.
