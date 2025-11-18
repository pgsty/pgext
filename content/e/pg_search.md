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
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/dev/pg_search" >}} | {{< ext "pg_search" >}} | `0.19.7` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} |

> [!Note] PG 17+ does not require dynamic loading, build by ParadeDB team 


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.19.7` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_search` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.19.7` | {{< bg "18" "pg_search_18" "red" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "green" >}} {{< bg "13" "pg_search_13" "red" >}} | `pg_search_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.19.7` | {{< bg "18" "postgresql-18-pg-search" "red" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "green" >}} {{< bg "13" "postgresql-13-pg-search" "red" >}} | `postgresql-$v-pg-search` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.3" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.3" "pg_search_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_search_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_search_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-search : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.19.7" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.19.7" "postgresql-14-pg-search : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-search : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.19.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.1 MiB | [pg_search_17-0.19.3-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.19.3-1PARADEDB.el8.x86_64.rpm) |
| `pg_search_17` | `0.19.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.5 MiB | [pg_search_17-0.19.3-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.19.3-1PARADEDB.el8.aarch64.rpm) |
| `pg_search_17` | `0.19.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.9 MiB | [pg_search_17-0.19.3-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.19.3-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_17` | `0.19.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.8 MiB | [pg_search_17-0.19.3-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.19.3-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.19.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.4 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.5 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.7 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.4 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 45.6 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-search` | `0.19.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.4 MiB | [postgresql-17-pg-search_0.19.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.19.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.19.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.1 MiB | [pg_search_16-0.19.3-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.19.3-1PARADEDB.el8.x86_64.rpm) |
| `pg_search_16` | `0.19.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.5 MiB | [pg_search_16-0.19.3-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.19.3-1PARADEDB.el8.aarch64.rpm) |
| `pg_search_16` | `0.19.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.9 MiB | [pg_search_16-0.19.3-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.19.3-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_16` | `0.19.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.8 MiB | [pg_search_16-0.19.3-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.19.3-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.19.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.5 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.5 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.7 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.4 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 45.6 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-search` | `0.19.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.3 MiB | [postgresql-16-pg-search_0.19.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.19.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.19.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.1 MiB | [pg_search_15-0.19.3-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.19.3-1PARADEDB.el8.x86_64.rpm) |
| `pg_search_15` | `0.19.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.5 MiB | [pg_search_15-0.19.3-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.19.3-1PARADEDB.el8.aarch64.rpm) |
| `pg_search_15` | `0.19.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.9 MiB | [pg_search_15-0.19.3-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.19.3-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_15` | `0.19.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.8 MiB | [pg_search_15-0.19.3-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.19.3-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.19.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.4 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.4 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.4 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 45.6 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-search` | `0.19.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.3 MiB | [postgresql-15-pg-search_0.19.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.19.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_14` | `0.19.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.1 MiB | [pg_search_14-0.19.3-1PARADEDB.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_14-0.19.3-1PARADEDB.el8.x86_64.rpm) |
| `pg_search_14` | `0.19.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 45.5 MiB | [pg_search_14-0.19.3-1PARADEDB.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_14-0.19.3-1PARADEDB.el8.aarch64.rpm) |
| `pg_search_14` | `0.19.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.9 MiB | [pg_search_14-0.19.3-1PARADEDB.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_14-0.19.3-1PARADEDB.el9.x86_64.rpm) |
| `pg_search_14` | `0.19.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.8 MiB | [pg_search_14-0.19.3-1PARADEDB.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_14-0.19.3-1PARADEDB.el9.aarch64.rpm) |
| `postgresql-14-pg-search` | `0.19.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.4 KiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.4 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.5 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.6 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.3 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 45.6 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-search` | `0.19.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.3 MiB | [postgresql-14-pg-search_0.19.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-14-pg-search_0.19.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/paradedb/paradedb/tree/dev/pg_search" title="Repository" icon="github" subtitle="github.com/paradedb/paradedb/tree/dev/pg_search" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.19.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_search;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_search;		# install via package name, for the active PG version

pig install pg_search -v 17;   # install for PG 17
pig install pg_search -v 16;   # install for PG 16
pig install pg_search -v 15;   # install for PG 15
pig install pg_search -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_search;
```


> [!NOTE] THIS EXTENSION is built by ParadeDB team and delivered by the PIGSTY repo

## Usage

https://docs.paradedb.com/documentation/getting-started/quickstart

```sql
CREATE EXTENSION pg_search;

ALTER SYSTEM SET paradedb.pg_search_telemetry TO 'off';

CALL paradedb.create_bm25_test_table(
  schema_name => 'public',
  table_name => 'mock_items'
);
    
SELECT description, rating, category FROM mock_items LIMIT 3;

CALL paradedb.create_bm25(
        index_name => 'search_idx',
        schema_name => 'public',
        table_name => 'mock_items',
        key_field => 'id',
        text_fields => paradedb.field('description', tokenizer => paradedb.tokenizer('en_stem')) ||
                       paradedb.field('category'),
        numeric_fields => paradedb.field('rating')
     );

SELECT description, rating, category
FROM search_idx.search('(description:keyboard OR category:electronics) AND rating:>2',limit_rows => 5);

CALL paradedb.create_bm25(
        index_name => 'ngrams_idx',
        schema_name => 'public',
        table_name => 'mock_items',
        key_field => 'id',
        text_fields => paradedb.field('description', tokenizer => paradedb.tokenizer('ngram', min_gram => 4, max_gram => 4, prefix_only => false)) ||
                       paradedb.field('category')
     );

SELECT description, rating, category
FROM ngrams_idx.search('description:blue');
```