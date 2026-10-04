---
title: "pg_search"
linkTitle: "pg_search"
description: "Full text search for PostgreSQL using BM25"
weight: 2100
categories: ["FTS"]
languages: ["Rust"]
licenses: ["AGPL-3.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_search**](https://github.com/paradedb/paradedb/tree/main/pg_search) : Full text search for PostgreSQL using BM25


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2100** | {{< badge content="pg_search" link="https://github.com/paradedb/paradedb/tree/main/pg_search" >}} | {{< ext "pg_search" >}} | `0.26.0` | {{< category "FTS" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `paradedb` |
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "pgfaceting" >}} {{< ext "roaringbitmap" >}} {{< ext "rum" >}} |

> [!Note] Requires preload and pgvector; conflicts with pg_textsearch and vchord_bm25. pgrx 0.19.3. Upgrading to 0.26.0 requires REINDEX of every vector-bearing index; recreate indexes with obsolete centroid_ratio/training_samples_per_centroid options.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_search` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.0` | {{< bg "18" "pg_search_18" "green" >}} {{< bg "17" "pg_search_17" "green" >}} {{< bg "16" "pg_search_16" "green" >}} {{< bg "15" "pg_search_15" "green" >}} {{< bg "14" "pg_search_14" "red" >}} | `pg_search_$v` | `pgvector_$v`, `openblas` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.0` | {{< bg "18" "postgresql-18-pg-search" "green" >}} {{< bg "17" "postgresql-17-pg-search" "green" >}} {{< bg "16" "postgresql-16-pg-search" "green" >}} {{< bg "15" "postgresql-15-pg-search" "green" >}} {{< bg "14" "postgresql-14-pg-search" "red" >}} | `postgresql-$v-pg-search` | `postgresql-$v-pgvector`, `libopenblas0` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "pg_search_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_search_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-18-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-17-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-16-pg-search : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.0" "postgresql-15-pg-search : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-search : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_18` | `0.26.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 74.8 MiB | [pg_search_18-0.26.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_18-0.26.0-1PGSTY.el8.x86_64.rpm) |
| `pg_search_18` | `0.26.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 71.5 MiB | [pg_search_18-0.26.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_18-0.26.0-1PGSTY.el8.aarch64.rpm) |
| `pg_search_18` | `0.26.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 76.9 MiB | [pg_search_18-0.26.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_18-0.26.0-1PGSTY.el9.x86_64.rpm) |
| `pg_search_18` | `0.26.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 75.3 MiB | [pg_search_18-0.26.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_18-0.26.0-1PGSTY.el9.aarch64.rpm) |
| `pg_search_18` | `0.26.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.6 MiB | [pg_search_18-0.26.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_18-0.26.0-1PGSTY.el10.x86_64.rpm) |
| `pg_search_18` | `0.26.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 75.2 MiB | [pg_search_18-0.26.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_18-0.26.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-search` | `0.26.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 64.4 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.5 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 64.4 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.7 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.7 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.1 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.5 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.1 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.6 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-search` | `0.26.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.0 MiB | [postgresql-18-pg-search_0.26.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-18-pg-search_0.26.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_17` | `0.26.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 74.9 MiB | [pg_search_17-0.26.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_17-0.26.0-1PGSTY.el8.x86_64.rpm) |
| `pg_search_17` | `0.26.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 71.5 MiB | [pg_search_17-0.26.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_17-0.26.0-1PGSTY.el8.aarch64.rpm) |
| `pg_search_17` | `0.26.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 77.0 MiB | [pg_search_17-0.26.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_17-0.26.0-1PGSTY.el9.x86_64.rpm) |
| `pg_search_17` | `0.26.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 75.3 MiB | [pg_search_17-0.26.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_17-0.26.0-1PGSTY.el9.aarch64.rpm) |
| `pg_search_17` | `0.26.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.6 MiB | [pg_search_17-0.26.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_17-0.26.0-1PGSTY.el10.x86_64.rpm) |
| `pg_search_17` | `0.26.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 75.1 MiB | [pg_search_17-0.26.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_17-0.26.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-search` | `0.26.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 64.5 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.7 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 64.3 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.6 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.7 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.0 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.5 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.8 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.6 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-search` | `0.26.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.1 MiB | [postgresql-17-pg-search_0.26.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-17-pg-search_0.26.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_16` | `0.26.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 74.9 MiB | [pg_search_16-0.26.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_16-0.26.0-1PGSTY.el8.x86_64.rpm) |
| `pg_search_16` | `0.26.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 71.5 MiB | [pg_search_16-0.26.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_16-0.26.0-1PGSTY.el8.aarch64.rpm) |
| `pg_search_16` | `0.26.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 77.0 MiB | [pg_search_16-0.26.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_16-0.26.0-1PGSTY.el9.x86_64.rpm) |
| `pg_search_16` | `0.26.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 75.3 MiB | [pg_search_16-0.26.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_16-0.26.0-1PGSTY.el9.aarch64.rpm) |
| `pg_search_16` | `0.26.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.7 MiB | [pg_search_16-0.26.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_16-0.26.0-1PGSTY.el10.x86_64.rpm) |
| `pg_search_16` | `0.26.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 75.1 MiB | [pg_search_16-0.26.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_16-0.26.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-search` | `0.26.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 64.4 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.7 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 64.4 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.5 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.7 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.0 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.5 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 73.8 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.7 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-search` | `0.26.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.1 MiB | [postgresql-16-pg-search_0.26.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-16-pg-search_0.26.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_search_15` | `0.26.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 74.8 MiB | [pg_search_15-0.26.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_search_15-0.26.0-1PGSTY.el8.x86_64.rpm) |
| `pg_search_15` | `0.26.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 71.4 MiB | [pg_search_15-0.26.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_search_15-0.26.0-1PGSTY.el8.aarch64.rpm) |
| `pg_search_15` | `0.26.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 76.9 MiB | [pg_search_15-0.26.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_search_15-0.26.0-1PGSTY.el9.x86_64.rpm) |
| `pg_search_15` | `0.26.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 75.1 MiB | [pg_search_15-0.26.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_search_15-0.26.0-1PGSTY.el9.aarch64.rpm) |
| `pg_search_15` | `0.26.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.6 MiB | [pg_search_15-0.26.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_search_15-0.26.0-1PGSTY.el10.x86_64.rpm) |
| `pg_search_15` | `0.26.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 75.2 MiB | [pg_search_15-0.26.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_search_15-0.26.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-search` | `0.26.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 64.3 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.6 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 64.3 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.6 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.8 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.2 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.5 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.1 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.6 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-search` | `0.26.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 74.0 MiB | [postgresql-15-pg-search_0.26.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-search/postgresql-15-pg-search_0.26.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/paradedb/paradedb/tree/main/pg_search" title="Repository" icon="github" subtitle="github.com/paradedb/paradedb/tree/main/pg_search" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_search-0.26.0.tar.gz pg_search_collect_third_party_licenses.py" />}}
{{< /cards >}}


```bash
pig build pkg pg_search;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

- [0.26.0 migration notes](https://github.com/paradedb/paradedb/blob/aa7f9dd407018036d144a54a3f875c5a00d20cda/docs/project/changelog/0.26.0.mdx)
- [0.26.0 SQL upgrade](https://github.com/paradedb/paradedb/blob/aa7f9dd407018036d144a54a3f875c5a00d20cda/pg_search/sql/pg_search--0.25.11--0.26.0.sql)
- [pg_search v0.25.11 README](https://github.com/paradedb/paradedb/blob/v0.25.11/pg_search/README.md)
- [pg_search v0.25.11 release](https://github.com/paradedb/paradedb/releases/tag/v0.25.11)
- [PGXN 0.25.11 metadata](https://api.pgxn.org/src/pg_search/pg_search-0.25.11/META.json)
- [pg_search v0.25.1 migration notes](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/changelog/0.25.1.mdx)
- [Create a ParadeDB index](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/indexing/create-index.mdx)
- [Full-text match operators](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/full-text/match.mdx)
- [BM25 scoring](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/sorting/score.mdx)
- [Highlighting and snippets](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/full-text/highlight.mdx)
- [Index vectors](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/indexing/indexing-vectors.mdx)
- [Query vectors](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/vector/querying.mdx)
- [Hybrid-search overview](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/hybrid/overview.mdx)

`pg_search` 0.26.0 adds ParadeDB's full-text, structured, vector, and hybrid search index to PostgreSQL. Version 0.25 uses the `paradedb` index access method; the older `bm25` access-method name remains a compatibility alias. The extension requires `vector`, supports PostgreSQL 15-18 upstream, and must be loaded through `shared_preload_libraries`.

### Install and Build an Index

```ini
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

### Version 0.25.11 and Caveats

- Version 0.25 renamed the primary index access method from `bm25` to `paradedb`. Existing `USING bm25` definitions remain supported, but new examples should use `USING paradedb`.
- Version 0.25.1 supports deterministic vector tie breakers and pushes the vector arm of reciprocal-rank-fusion queries into the index. It also adds `paradedb.vector_clustering_threshold`, whose default is 500, and caps vector-index build parallelism at four workers.
- Version 0.25.1 removes `paradedb.vector_cluster_probe_epsilon` and changes the vector-index bounds gate. After upgrading a database from 0.25.0, `REINDEX` every ParadeDB index that contains a vector field; installing the new shared library and running `ALTER EXTENSION` alone is not sufficient for those indexes.
- Version 0.25.2 is a stability and correctness release. It fixes fieldless `more_like_this` with vector columns, `pdb.fuzzy` under generic prepared plans, orphaned dynamic filters, several parallel subplan and MPP plan-shape errors, and tightens access controls for typemod definitions. It adds no further index migration beyond the inherited 0.25.0 vector-index rebuild.
- Versions 0.25.4 through 0.25.6 add `paradedb.vector_clusters`, partition-aware index builds, aggregate score joins, unified planner-warning controls, and bitmap-scan intersection. They also fix dropped bitmap-intersection children and sortable encoding for negative high-precision `numeric` values. Validate plans and ordering after the upgrade even when no explicit new index migration is documented.
- Versions 0.25.7–0.25.10 extend partition-aware concurrent builds, lateral array unnest and join aggregates, and parallel aggregation. Version 0.25.10 restores parallel top-K planning and fixes HOT-redirect row loss, missing JSON-path values, and aggregate defaults whose types cannot preserve the requested value. Review result correctness and query plans after upgrading.
- Version 0.25.11 fixes a backend crash when a search scan is cancelled or terminated and treats a null `STRING_AGG` delimiter as an empty separator.
- `CREATE EXTENSION pg_search CASCADE` can install the required `vector` extension, but every server process still needs the preload configuration and restart first. Loading it only with `LOAD` or `session_preload_libraries` is insufficient.
- Query plans, tokenization, and ranking can change when an index is rebuilt with different field options. Test relevance and vector recall with production-shaped data before rollout.

### 0.26.0 Migration

Version 0.26.0 changes vector storage to format 4. After updating the library and SQL extension, rebuild every index containing a vector field, including indexes with quantization disabled. Until rebuilt, vector queries fail with a rebuild-required error; writes continue and `REINDEX CONCURRENTLY` is supported. Quantization defaults on for vectors with at least 64 dimensions. Its settings take effect only during index creation or rebuilding. Indexes storing `centroid_ratio` or `training_samples_per_centroid` must be recreated using `training_sample_ratio` and `max_leaf_size`. Rebuild indexes with multiple regex-tokenized fields to correct pattern selection as well.

The release adds beta `partition_by` index partitioning, window-aggregate pushdown over joins, and opt-in `paradedb.spill_to_disk`. Aggregate `visibility` defaults to `transaction`; `raw` skips MVCC checks and must not be mistaken for transaction-visible results. Recheck ranking, recall, plans and storage requirements before application rollout.
