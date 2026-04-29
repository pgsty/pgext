---
title: "pg_tokenizer"
linkTitle: "pg_tokenizer"
description: "Tokenizers for full-text search"
weight: 2160
categories: ["FTS"]
width: full
---

[**pg_tokenizer**](https://github.com/tensorchord/pg_tokenizer.rs) : Tokenizers for full-text search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2160** | {{< badge content="pg_tokenizer" link="https://github.com/tensorchord/pg_tokenizer.rs" >}} | {{< ext "pg_tokenizer" >}} | `0.1.1` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tokenizer_catalog` |
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_trgm" >}} |

> [!Note] PG18 fix by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_tokenizer` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "pg_tokenizer_18" "green" >}} {{< bg "17" "pg_tokenizer_17" "green" >}} {{< bg "16" "pg_tokenizer_16" "green" >}} {{< bg "15" "pg_tokenizer_15" "green" >}} {{< bg "14" "pg_tokenizer_14" "green" >}} | `pg_tokenizer_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pg-tokenizer" "green" >}} {{< bg "17" "postgresql-17-pg-tokenizer" "green" >}} {{< bg "16" "postgresql-16-pg-tokenizer" "green" >}} {{< bg "15" "postgresql-15-pg-tokenizer" "green" >}} {{< bg "14" "postgresql-14-pg-tokenizer" "green" >}} | `postgresql-$v-pg-tokenizer` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tokenizer : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tokenizer : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tokenizer : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.0 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.9 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.0 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.7 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.6 MiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.0 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.9 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.0 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_16` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.0 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.9 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.0 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_15` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.0 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.9 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.0 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_14` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.0 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.9 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.0 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/pg_tokenizer.rs" title="Repository" icon="github" subtitle="github.com/tensorchord/pg_tokenizer.rs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tokenizer.rs-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tokenizer;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tokenizer;		# install via package name, for the active PG version

pig install pg_tokenizer -v 18;   # install for PG 18
pig install pg_tokenizer -v 17;   # install for PG 17
pig install pg_tokenizer -v 16;   # install for PG 16
pig install pg_tokenizer -v 15;   # install for PG 15
pig install pg_tokenizer -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_tokenizer';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tokenizer;
```



## Usage

> [GitHub: tensorchord/pg_tokenizer.rs](https://github.com/tensorchord/pg_tokenizer.rs)

`pg_tokenizer` is a PostgreSQL extension that provides tokenizers for full-text search. It is designed to work with [VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25) for native BM25 ranking index support.

## Quick Start

```sql
CREATE EXTENSION pg_tokenizer;

-- Create a tokenizer using the LLMLingua2 model
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- Tokenize text
SELECT tokenize('PostgreSQL is a powerful, open-source object-relational database system. It has over 15 years of active development.', 'tokenizer1');
```

## Tokenizer Models

pg_tokenizer supports multiple tokenizer models for different languages and use cases:

| Model | Language | Description |
|-------|----------|-------------|
| `llmlingua2` | English | BERT-based tokenizer from LLMLingua2 |
| `jieba` | Chinese | Jieba Chinese text segmentation |
| `lindera/ipadic` | Japanese | Lindera tokenizer with IPADIC dictionary |
| Custom models | Any | User-trained models for domain-specific text |

### Creating Tokenizers

```sql
-- English tokenizer
SELECT create_tokenizer('en_tokenizer', $$
model = "llmlingua2"
$$);

-- Chinese tokenizer
SELECT create_tokenizer('zh_tokenizer', $$
model = "jieba"
$$);

-- Japanese tokenizer
SELECT create_tokenizer('ja_tokenizer', $$
model = "lindera/ipadic"
$$);
```

### Tokenizing Text

```sql
-- Tokenize English text
SELECT tokenize('full text search in PostgreSQL', 'en_tokenizer');

-- Tokenize Chinese text
SELECT tokenize('PostgreSQL是一个强大的数据库系统', 'zh_tokenizer');
```

## Text Analyzer

pg_tokenizer also provides text analyzer functionality that combines tokenization with additional text processing steps. For detailed text analyzer usage, refer to the [Text Analyzer documentation](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md).

## Integration with VectorChord-BM25

pg_tokenizer is typically used together with VectorChord-BM25 for full BM25 ranking support:

```sql
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- Create a tokenizer
SELECT create_tokenizer('my_tokenizer', $$
model = "llmlingua2"
$$);

-- Tokenize text into bm25vectors for indexing and search
SELECT tokenize('your search query', 'my_tokenizer');
```

## Documentation

For more details, see the full documentation:

- [Installation](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/01-installation.md)
- [Examples](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/03-examples.md)
- [Usage](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/04-usage.md)
- [Text Analyzer](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md)
- [Model Reference](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/06-model.md)
- [Limitations](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/07-limitation.md)
