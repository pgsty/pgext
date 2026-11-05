---
title: "pg_tokenizer"
linkTitle: "pg_tokenizer"
description: "Tokenizers for full-text search"
weight: 2170
categories: ["FTS"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_tokenizer**](https://vchord.pgsty.com/docs/tokenizer) : Tokenizers for full-text search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2170** | {{< badge content="pg_tokenizer" link="https://vchord.pgsty.com/docs/tokenizer" >}} | {{< ext "pg_tokenizer" >}} | `1.0.0` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tokenizer_catalog` |
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pg_jieba" >}} {{< ext "pg_cjk_parser" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg_tiktoken_c" >}} {{< ext "unaccent" >}} {{< ext "dict_xsyn" >}} {{< ext "dict_int" >}} {{< ext "hunspell_cs_cz" >}} {{< ext "pg_kazsearch" >}} |

> [!Note] PGSTY fork 1.0.0; native tsvector; preload required. Preserves catalog data; Lindera 6 needs matching external dictionaries. Recompute stored tsvectors for changed token IDs/phrase positions.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_tokenizer` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_tokenizer_18" "green" >}} {{< bg "17" "pg_tokenizer_17" "green" >}} {{< bg "16" "pg_tokenizer_16" "green" >}} {{< bg "15" "pg_tokenizer_15" "green" >}} {{< bg "14" "pg_tokenizer_14" "green" >}} | `pg_tokenizer_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-tokenizer" "green" >}} {{< bg "17" "postgresql-17-pg-tokenizer" "green" >}} {{< bg "16" "postgresql-16-pg-tokenizer" "green" >}} {{< bg "15" "postgresql-15-pg-tokenizer" "green" >}} {{< bg "14" "postgresql-14-pg-tokenizer" "green" >}} | `postgresql-$v-pg-tokenizer` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_tokenizer_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.1 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_18-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_18-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.2 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_18-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.9 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_18-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_18-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 MiB | [pg_tokenizer_18-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_18-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.9 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.9 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.6 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.6 MiB | [postgresql-18-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.1 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_17-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_17-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.2 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_17-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.9 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_17-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_17-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 MiB | [pg_tokenizer_17-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_17-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.9 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.9 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.6 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.6 MiB | [postgresql-17-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_16-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_16-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.2 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_16-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.9 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_16-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_16-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 MiB | [pg_tokenizer_16-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_16-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.9 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.9 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.6 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.6 MiB | [postgresql-16-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_15-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_15-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.2 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_15-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.9 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_15-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_15-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 MiB | [pg_tokenizer_15-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_15-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.3 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.9 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.9 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.6 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.8 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.6 MiB | [postgresql-15-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_14-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.4 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_14-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_14-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.9 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_14-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.1 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_14-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 MiB | [pg_tokenizer_14-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_14-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.9 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.9 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.6 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.8 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.6 MiB | [postgresql-14-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://vchord.pgsty.com/docs/tokenizer" title="Repository" icon="link" subtitle="vchord.pgsty.com/docs/tokenizer" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tokenizer-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_tokenizer;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_tokenizer;		# install via package name, for the active PG version

pig install pg_tokenizer -v 18;   # install for PG 18
pig install pg_tokenizer -v 17;   # install for PG 17
pig install pg_tokenizer -v 16;   # install for PG 16
pig install pg_tokenizer -v 15;   # install for PG 15
pig install pg_tokenizer -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_tokenizer';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_tokenizer;
```

## Usage

Sources:

- [Installation, upgrades, permissions, and backups](https://vchord.pgsty.com/docs/tokenizer/installation/)
- [Tokenizer and native full-text workflows](https://vchord.pgsty.com/docs/tokenizer/overview/)
- [Text analyzer pipelines](https://vchord.pgsty.com/docs/tokenizer/text-analyzer/)
- [Models and filtered corpus training](https://vchord.pgsty.com/docs/tokenizer/models/)
- [Multilingual tokenization](https://vchord.pgsty.com/docs/tokenizer/multilingual/)
- [SQL and TOML reference](https://vchord.pgsty.com/docs/tokenizer/reference/)
- [1.0.0 control file](https://github.com/pgsty/pg_tokenizer/blob/v1.0.0/pg_tokenizer.control)

`pg_tokenizer` 1.0.0 provides configurable text analysis and model tokenization on PostgreSQL 14–18. It returns string tokens, integer token IDs, or native `tsvector` values for GIN/GiST and compatible BM25 indexes. It has no required extension dependency, needs shared preloading, and installs into the fixed `tokenizer_catalog` schema.

### Analyze and Index Text

Have an administrator add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_text_analyzer('fts_words', $$
character_filters = ["to_lowercase"]
pre_tokenizer = "unicode_segmentation"
$$);
SELECT apply_text_analyzer('Cat dog cat', 'fts_words');
SELECT analyze_to_tsvector('Cat dog cat', 'fts_words');

CREATE TABLE fts_documents (id bigint PRIMARY KEY, passage text, terms tsvector);
INSERT INTO fts_documents VALUES
    (1, 'Cat dog cat', analyze_to_tsvector('Cat dog cat', 'fts_words'));
CREATE INDEX fts_documents_gin ON fts_documents USING gin (terms);
SELECT id FROM fts_documents WHERE terms @@ 'cat <-> dog'::tsquery;
```

`apply_text_analyzer` returns `text[]`; `analyze_to_tsvector` uses those strings as lexemes with positions. Analyzer definitions are catalog data, so these functions are `STABLE`. Store their output in an ordinary column and maintain it on writes; do not use them directly in an expression index or stored generated column.

```sql
CREATE FUNCTION fts_documents_set_terms() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
    NEW.terms := tokenizer_catalog.analyze_to_tsvector(NEW.passage, 'fts_words');
    RETURN NEW;
END;
$$;
CREATE TRIGGER fts_documents_terms
BEFORE INSERT OR UPDATE OF passage ON fts_documents
FOR EACH ROW EXECUTE FUNCTION fts_documents_set_terms();
```

The phrase example uses known analyzer lexemes. The extension does not supply a general `tsquery` builder: query terms must use the same analysis and representation as stored documents. Positions describe the final token sequence, not character offsets in the source text. Filters can change phrase distances. PostgreSQL limits still apply: long lexemes are skipped, positions above 16,383 are clamped, and each lexeme retains at most 256 positions.

### Chinese Text and Model Tokenization

Jieba is an analyzer pre-tokenizer. It can produce native Chinese lexemes without training a model:

```sql
SELECT create_text_analyzer('zh_words', $$
[pre_tokenizer.jieba]
mode = "search"
enable_hmm = true
$$);
SELECT analyze_to_tsvector('PostgreSQL全文检索', 'zh_words');

SELECT create_tokenizer('bert_tokens', $$
model = "bert_base_uncased"
$$);
SELECT tokenize('hello hello', 'bert_tokens');
SELECT tokenize_to_tsvector('hello hello', 'bert_tokens');
```

`tokenize` returns `integer[]`; `tokenize_to_tsvector` stores the model IDs as decimal lexemes. The two native vector functions preserve repeated output tokens without decoding the model or applying another PostgreSQL normalization step. A normal language query such as `plainto_tsquery` does not reproduce model IDs. Use matching vectors with the BM25 query constructor, or construct a native query from the matching lexemes.

### Function and Configuration Index

| Interface | Purpose |
| --- | --- |
| `create_text_analyzer`, `drop_text_analyzer` | Manage TOML pipelines of character filters, a pre-tokenizer, and token filters. |
| `create_tokenizer`, `drop_tokenizer` | Combine an analyzer with a vocabulary or subword model. |
| `create_custom_model_tokenizer_and_trigger` | Train a custom vocabulary, backfill integer token IDs, and create maintenance triggers; an optional seventh `filter` argument selects corpus rows. |
| `create_huggingface_model`, `create_lindera_model` | Import a Hugging Face tokenizer definition or configure a Lindera dictionary. |
| `create_stopwords`, `create_synonym` | Create named dictionaries for analyzer filters. |
| `add_preload_model`, `remove_preload_model`, `list_preload_models` | Manage or inspect model preloading; changes require authorized administrators and a restart. |
| `llmlingua2`, `bert_base_uncased`, `wiki_tocken`, `gemma2b` | Built-in model names. |

The custom-model helper still fills integer arrays, not native vectors. Its filter is a trusted SQL Boolean expression over source columns: only true rows participate, the vocabulary remains append-only, and rows leaving the filter keep their previous target value. Do not pass untrusted SQL. Keep a separate vector-maintenance step when using native full-text indexes.

### Upgrade, Permissions, and Backups

Install the 1.0.0 files and restart PostgreSQL before updating each database:

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '1.0.0';
```

Upgrade paths from 0.1.0 and 0.1.1 preserve catalog rows, custom-model IDs, and the six-argument helper. The upgrade refuses to proceed while expression indexes, generated columns, partition keys, or extended statistics depend on the old immutable analyzer function; replace these uses with maintained columns first. Committed definition changes now invalidate backend caches according to SQL snapshots, but never rewrite stored token arrays or vectors.

Lindera 6 changes dictionary token IDs. Rebuild external dictionaries with Lindera 6, regenerate affected stored tokens/vectors, and rebuild indexes. Native RPM/DEB builds do not embed Lindera dictionaries; provide the external dictionary yourself. Back up those files and the cluster's `pg_tokenizer/preload_models` directory separately from a logical dump.

`pg_database_owner` receives access to database-local catalogs and functions; application roles need explicit grants. Object ownership still controls removal. Cluster preload changes are nontransactional and restricted separately. Vocabulary writers introducing new tokens serialize per model; retry SQLSTATE `40001` under repeatable-read or serializable isolation. Definition-changing APIs disallow two-phase prepare.

An ordinary `pg_dump` now includes analyzer, tokenizer, model, stopword, and synonym definitions, together with custom vocabulary tables, sequences, and source triggers. The shipped `lucene_english`, `nltk_english`, and `iso_english` lists are reserved and recreated by extension installation; use another name for a custom list.
