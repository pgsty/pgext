---
title: "pg_textsearch"
linkTitle: "pg_textsearch"
description: "Full-text search with BM25 ranking"
weight: 2180
categories: ["FTS"]
width: full
---

[**pg_textsearch**](https://github.com/timescale/pg_textsearch) : Full-text search with BM25 ranking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2180** | {{< badge content="pg_textsearch" link="https://github.com/timescale/pg_textsearch" >}} | {{< ext "pg_textsearch" >}} | `1.1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "biscuit" >}} {{< ext "fuzzystrmatch" >}} |

> [!Note] bm25 am conflicts with pg_search; must be preloaded via shared_preload_libraries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_textsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_textsearch_18" "green" >}} {{< bg "17" "pg_textsearch_17" "green" >}} {{< bg "16" "pg_textsearch_16" "red" >}} {{< bg "15" "pg_textsearch_15" "red" >}} {{< bg "14" "pg_textsearch_14" "red" >}} | `pg_textsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-textsearch" "green" >}} {{< bg "17" "postgresql-17-textsearch" "green" >}} {{< bg "16" "postgresql-16-textsearch" "red" >}} {{< bg "15" "postgresql-15-textsearch" "red" >}} {{< bg "14" "postgresql-14-textsearch" "red" >}} | `postgresql-$v-textsearch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 121.4 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 115.5 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 112.5 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 109.7 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.5 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 111.7 KiB | [pg_textsearch_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-textsearch` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1003.2 KiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 994.4 KiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1003.6 KiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 995.5 KiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.0 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.0 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.1 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-textsearch` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-18-textsearch_1.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 121.3 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 115.3 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 112.3 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 109.5 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.1 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 111.8 KiB | [pg_textsearch_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-textsearch` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 990.5 KiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 980.4 KiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 991.9 KiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 982.3 KiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-textsearch` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_textsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_textsearch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_textsearch-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_textsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_textsearch;		# install via package name, for the active PG version

pig install pg_textsearch -v 18;   # install for PG 18
pig install pg_textsearch -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_textsearch';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_textsearch;
```

## Usage

Sources: [README v1.1.0](https://github.com/timescale/pg_textsearch/blob/v1.1.0/README.md), [v1.1.0 release notes](https://github.com/timescale/pg_textsearch/releases/tag/v1.1.0)

`pg_textsearch` provides BM25-ranked full-text search for PostgreSQL with a `bm25` access method and the `<@>` scoring operator. Upstream marks `v1.1.0` as production ready.

`v1.1.0` supports PostgreSQL 17 and 18. Prebuilt release assets are published for both PostgreSQL versions on Linux and macOS. The extension must be loaded through `shared_preload_libraries` before `CREATE EXTENSION`.

### Enable the Extension

```ini
shared_preload_libraries = 'pg_textsearch'  # add to any existing list
```

```sql
CREATE EXTENSION pg_textsearch;
```

Install the new binary and restart PostgreSQL before running an extension upgrade:

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

Upstream says upgrading from 1.0.0 to 1.1.0 does not require `REINDEX`.

### Build and Query BM25 Indexes

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);

CREATE INDEX docs_idx
ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT *
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

`<@>` returns the negative BM25 score because PostgreSQL operator index scans are ascending; lower values are better matches. Use `ORDER BY ... LIMIT` for fast top-k searches.

For an explicit index reference, use `to_bm25query()`:

```sql
SELECT *
FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

The main documented SQL surface is:

- `text <@> 'query'` to score text with planner-detected index context.
- `text <@> bm25query` to score with an explicit `bm25query`.
- `to_bm25query(text)` for ORDER BY use with planner-selected index context.
- `to_bm25query(text, text)` for query text plus index name.
- `bm25query = bm25query` for equality checks.

### Index Options and Data Shapes

```sql
CREATE INDEX ON documents USING bm25(content)
WITH (text_config = 'english', k1 = 1.5, b = 0.8);
```

Index options are `text_config` (required), `k1` (default `1.2`), and `b` (default `0.75`). Text search configurations such as `english`, `simple`, `french`, and `german` use PostgreSQL text search configuration names.

`v1.1.0` adds native array input support for `text[]`, `varchar[]`, and `bpchar[]` columns; array elements are concatenated before tokenization.

```sql
CREATE TABLE posts (id serial PRIMARY KEY, tags text[]);
CREATE INDEX posts_tags_idx ON posts USING bm25(tags)
WITH (text_config = 'english');

SELECT *
FROM posts
ORDER BY tags <@> 'database'
LIMIT 10;
```

Expression indexes support immutable text expressions, including JSONB extraction, text transformations, and multi-column concatenation:

```sql
CREATE INDEX events_msg_idx ON events USING bm25 ((data->>'message'))
WITH (text_config = 'english');

SELECT *
FROM events
ORDER BY (data->>'message') <@> to_bm25query('network error', 'events_msg_idx')
LIMIT 10;
```

Partial indexes scope search to a subset of rows. Query them with an explicit index name:

```sql
CREATE INDEX docs_en_idx ON docs USING bm25(content)
WITH (text_config = 'english')
WHERE lang = 'en';

SELECT *
FROM docs
WHERE lang = 'en'
ORDER BY content <@> to_bm25query('databases', 'docs_en_idx')
LIMIT 10;
```

### Operations and GUCs

```sql
SELECT bm25_force_merge('docs_idx');
SELECT * FROM bm25_memory_usage();
```

`bm25_force_merge(index_name)` consolidates all segments into one and is best used after bulk loads, not during steady write traffic. `bm25_memory_usage()` reports shared memory usage for memtables.

Documented `pg_textsearch` GUCs in v1.1.0 include:

- `pg_textsearch.default_limit`
- `pg_textsearch.compress_segments`
- `pg_textsearch.segments_per_level`
- `pg_textsearch.memory_limit`
- `pg_textsearch.bulk_load_threshold`
- `pg_textsearch.memtable_spill_threshold` (deprecated; use `memory_limit` for new deployments)

`pg_textsearch.memory_limit` defaults to `2GB` and caps dynamic shared memory used by memtables. The release notes also call out improved concurrent insert throughput, faster VACUUM via segment alive bitsets, subtransaction cleanup, and parallel build race fixes.

### Caveats

- `pg_textsearch` requires `shared_preload_libraries = 'pg_textsearch'` and a PostgreSQL restart before `CREATE EXTENSION`.
- Inside PL/pgSQL and stored procedures, the implicit `text <@> 'query'` form does not use planner hooks; upstream says to use `to_bm25query()` with an explicit index name there.
- Phrase queries are not native because the index stores term frequencies, not term positions; use BM25 ranking plus a post-filter for phrase-like matching.
- Partial indexes require `to_bm25query()` with the index name because the implicit query form skips them.
- BM25 indexes on partitioned tables use partition-local statistics, so cross-partition scores may not be directly comparable.
- Words longer than PostgreSQL's `tsvector` word length limit are ignored during tokenization.
