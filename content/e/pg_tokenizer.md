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

[**pg_tokenizer**](https://github.com/supervc-stack/pg_tokenizer.rs) : Tokenizers for full-text search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2170** | {{< badge content="pg_tokenizer" link="https://github.com/supervc-stack/pg_tokenizer.rs" >}} | {{< ext "pg_tokenizer" >}} | `0.1.1` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tokenizer_catalog` |
|   **See Also**    | {{< ext "pgroonga" >}} {{< ext "pg_jieba" >}} {{< ext "pg_cjk_parser" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg_tiktoken_c" >}} {{< ext "unaccent" >}} {{< ext "dict_xsyn" >}} {{< ext "dict_int" >}} {{< ext "hunspell_cs_cz" >}} {{< ext "pg_kazsearch" >}} |

> [!Note] PG18 fix by Vonng.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_tokenizer` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "pg_tokenizer_18" "green" >}} {{< bg "17" "pg_tokenizer_17" "green" >}} {{< bg "16" "pg_tokenizer_16" "green" >}} {{< bg "15" "pg_tokenizer_15" "green" >}} {{< bg "14" "pg_tokenizer_14" "green" >}} | `pg_tokenizer_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pg-tokenizer" "green" >}} {{< bg "17" "postgresql-17-pg-tokenizer" "green" >}} {{< bg "16" "postgresql-16-pg-tokenizer" "green" >}} {{< bg "15" "postgresql-15-pg-tokenizer" "green" >}} {{< bg "14" "postgresql-14-pg-tokenizer" "green" >}} | `postgresql-$v-pg-tokenizer` | - |
{.packages}


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
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_18-0.1.1-3PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.0 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_18-0.1.1-3PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.5 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_18-0.1.1-3PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_18-0.1.1-3PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_18-0.1.1-3PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.3 MiB | [pg_tokenizer_18-0.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_18-0.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.1 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.9 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.1 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.9 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.2 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.1 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-tokenizer` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 MiB | [postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_17-0.1.1-3PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.0 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_17-0.1.1-3PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.5 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_17-0.1.1-3PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_17-0.1.1-3PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_17-0.1.1-3PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.3 MiB | [pg_tokenizer_17-0.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_17-0.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.1 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.1 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.2 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.1 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-tokenizer` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 MiB | [postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_16` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_16-0.1.1-3PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.0 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_16-0.1.1-3PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.5 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_16-0.1.1-3PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_16-0.1.1-3PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_16-0.1.1-3PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.3 MiB | [pg_tokenizer_16-0.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_16-0.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.1 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.1 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.2 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.1 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-tokenizer` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 MiB | [postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_15` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_15-0.1.1-3PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.0 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_15-0.1.1-3PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.5 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_15-0.1.1-3PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_15-0.1.1-3PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.3 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_15-0.1.1-3PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_15` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.3 MiB | [pg_tokenizer_15-0.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_15-0.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.1 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.1 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.1 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.1 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-tokenizer` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 MiB | [postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_14` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_14-0.1.1-3PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.0 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_14-0.1.1-3PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_14-0.1.1-3PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_14-0.1.1-3PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.3 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_14-0.1.1-3PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_14` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.3 MiB | [pg_tokenizer_14-0.1.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_14-0.1.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.1 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.1 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.2 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.1 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.0 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.1 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-tokenizer` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.0 MiB | [postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-3PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supervc-stack/pg_tokenizer.rs" title="Repository" icon="github" subtitle="github.com/supervc-stack/pg_tokenizer.rs" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tokenizer.rs-0.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_tokenizer;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [0.1.1 control](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/pg_tokenizer.control)
- [Installation and preload](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer usage](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/04-usage.md)
- [Reference and n-grams](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/00-reference.md)
- [Models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)
- [Cache limitations](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/07-limitation.md)

`pg_tokenizer` converts text to token IDs for search applications, commonly with VectorChord-BM25. A tokenizer combines text analysis with a vocabulary model. It requires shared preloading and installs its SQL objects in the fixed `tokenizer_catalog` schema.

### Enable and Tokenize

Add the library to the existing preload list and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_tokenizer('english', $$
model = "llmlingua2"
$$);
SELECT tokenize('PostgreSQL full text search', 'english');
```

`tokenize(text, text)` returns integer token IDs, not a relevance score or one row per token. With the BM25 companion installed, the returned array can be converted to its sparse vector type.

### Analyze Chinese Text

Jieba is a **pre-tokenizer**, not a built-in vocabulary model named jieba. Build a text analyzer and a custom model for the corpus:

```sql
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    token_ids integer[]
);
SELECT create_text_analyzer('chinese', $$
[pre_tokenizer.jieba]
$$);
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'chinese',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'token_ids'
);
INSERT INTO documents(passage) VALUES ('PostgreSQL全文检索');
SELECT tokenize('数据库', 'zh_tokenizer');
```

The helper learns a vocabulary from the source column and creates a trigger to maintain token IDs. Use the same tokenizer/model for documents and queries. Japanese uses an explicitly created Lindera model with a configured dictionary; a bare model name from another tokenizer API is not interchangeable.

### Object and Configuration Index

- `create_tokenizer`, `drop_tokenizer`, `tokenize`: manage and execute tokenizers.
- `create_text_analyzer`, `apply_text_analyzer`: run character filters, pre-tokenization, and token filters.
- `create_custom_model_tokenizer_and_trigger`, `create_lindera_model`, `create_huggingface_model`: create corpus or imported models.
- `create_stopwords`, `create_synonym`: manage dictionaries.
- Built-in models include `llmlingua2`, `bert_base_uncased`, `wiki_tocken`, and `gemma2b`.
- Version 0.1.1 adds the `ngram` token filter; `min_gram` and `max_gram` range from 1 to 255, while `preserve_original` defaults to false. Configurations use TOML.

### Upgrade and Cache Boundaries

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '0.1.1';
```

After replacing a preloaded library, restart PostgreSQL before updating database objects. The 0.1.0-to-0.1.1 migration adds no SQL objects; behavior changes are in the library. Analyzer, model, and tokenizer objects are cached per connection, and this cache does not follow transaction isolation or rollback. Reconnect or use the relevant drop helper to clear an object retained after rollback. The extension is not relocatable after creation. It supplies tokenization; ranking and search indexes belong to its consumers.
