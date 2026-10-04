---
title: "psql_bm25s"
linkTitle: "psql_bm25s"
description: "PostgreSQL extension for BM25-family lexical retrieval"
weight: 2220
categories: ["FTS"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**psql_bm25s**](https://github.com/Intelligent-Internet/Evoke) : PostgreSQL extension for BM25-family lexical retrieval


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2220** | {{< badge content="psql_bm25s" link="https://github.com/Intelligent-Internet/Evoke" >}} | {{< ext "psql_bm25s" >}} | `0.4.14` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_fts" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_rrf" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] Supports PostgreSQL 17-18; optional shared_preload_libraries arena is not required for normal use.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.14` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `psql_bm25s` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.14` | {{< bg "18" "psql_bm25s_18" "green" >}} {{< bg "17" "psql_bm25s_17" "green" >}} {{< bg "16" "psql_bm25s_16" "red" >}} {{< bg "15" "psql_bm25s_15" "red" >}} {{< bg "14" "psql_bm25s_14" "red" >}} | `psql_bm25s_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.14` | {{< bg "18" "postgresql-18-psql-bm25s" "green" >}} {{< bg "17" "postgresql-17-psql-bm25s" "green" >}} {{< bg "16" "postgresql-16-psql-bm25s" "red" >}} {{< bg "15" "postgresql-15-psql-bm25s" "red" >}} {{< bg "14" "postgresql-14-psql-bm25s" "red" >}} | `postgresql-$v-psql-bm25s` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "psql_bm25s_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "psql_bm25s_16 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_15 : N/A 0" "gray" >}} | {{< bg "N/A" "psql_bm25s_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.14" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-psql-bm25s : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-psql-bm25s : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `psql_bm25s_18` | `0.4.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 502.8 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/psql_bm25s_18-0.4.14-1PGSTY.el8.x86_64.rpm) |
| `psql_bm25s_18` | `0.4.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 484.6 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/psql_bm25s_18-0.4.14-1PGSTY.el8.aarch64.rpm) |
| `psql_bm25s_18` | `0.4.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 500.5 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/psql_bm25s_18-0.4.14-1PGSTY.el9.x86_64.rpm) |
| `psql_bm25s_18` | `0.4.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 485.3 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/psql_bm25s_18-0.4.14-1PGSTY.el9.aarch64.rpm) |
| `psql_bm25s_18` | `0.4.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 507.5 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/psql_bm25s_18-0.4.14-1PGSTY.el10.x86_64.rpm) |
| `psql_bm25s_18` | `0.4.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 490.9 KiB | [psql_bm25s_18-0.4.14-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/psql_bm25s_18-0.4.14-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 497.4 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 476.3 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 499.4 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 479.2 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 527.1 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 511.6 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 510.3 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~noble_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 497.3 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~noble_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u26.x86_64](/os/u26.x86_64) | pigsty | 506.9 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.14` | [u26.aarch64](/os/u26.aarch64) | pigsty | 492.8 KiB | [postgresql-18-psql-bm25s_0.4.14-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.14-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `psql_bm25s_17` | `0.4.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 502.7 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/psql_bm25s_17-0.4.14-1PGSTY.el8.x86_64.rpm) |
| `psql_bm25s_17` | `0.4.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 484.7 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/psql_bm25s_17-0.4.14-1PGSTY.el8.aarch64.rpm) |
| `psql_bm25s_17` | `0.4.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 500.3 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/psql_bm25s_17-0.4.14-1PGSTY.el9.x86_64.rpm) |
| `psql_bm25s_17` | `0.4.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 485.2 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/psql_bm25s_17-0.4.14-1PGSTY.el9.aarch64.rpm) |
| `psql_bm25s_17` | `0.4.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 507.4 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/psql_bm25s_17-0.4.14-1PGSTY.el10.x86_64.rpm) |
| `psql_bm25s_17` | `0.4.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 490.9 KiB | [psql_bm25s_17-0.4.14-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/psql_bm25s_17-0.4.14-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 497.5 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 475.6 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 499.5 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 479.1 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 553.4 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 538.1 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 510.2 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~noble_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 497.4 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~noble_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u26.x86_64](/os/u26.x86_64) | pigsty | 506.9 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.14` | [u26.aarch64](/os/u26.aarch64) | pigsty | 492.8 KiB | [postgresql-17-psql-bm25s_0.4.14-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.14-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Intelligent-Internet/Evoke" title="Repository" icon="github" subtitle="github.com/Intelligent-Internet/Evoke" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="psql_bm25s-0.4.14.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg psql_bm25s;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install psql_bm25s;		# install via package name, for the active PG version

pig install psql_bm25s -v 18;   # install for PG 18
pig install psql_bm25s -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION psql_bm25s;
```

## Usage

Sources:

- [README v0.4.14](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/README.md)
- [API reference](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/docs/api-reference.md)
- [query semantics](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/docs/query-semantics.md)
- [input types](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/docs/input-types.md)
- [index parameters](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/docs/index-parameters.md)
- [index policy](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/docs/index-policy.md)
- [Control file](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/psql_bm25s.control)
- [Upgrade SQL](https://github.com/Intelligent-Internet/Evoke/blob/v0.4.14/sql/psql_bm25s--0.4.11--0.4.14.sql)

`psql_bm25s` is a PostgreSQL-native index access method for BM25-family lexical retrieval. It keeps the BM25 contract explicit through corpus-statistics-driven ranking, exact top-k retrieval APIs, and PostgreSQL storage/maintenance behavior for mutable tables.

Version `0.4.14` is packaged for PostgreSQL 17 and 18 in this catalog.

### Basic Search

```sql
CREATE EXTENSION psql_bm25s;

CREATE TABLE docs (
    id integer PRIMARY KEY,
    title text NOT NULL,
    body text NOT NULL
);

INSERT INTO docs (id, title, body) VALUES
    (1, 'red apple', 'fresh red apple fruit'),
    (2, 'green apple', 'green apple slices'),
    (3, 'orange citrus', 'orange citrus fruit'),
    (4, 'cat guide', 'small cat animal care');

CREATE INDEX docs_bm25_idx
    ON docs USING psql_bm25s (body);

SELECT d.id, d.title, h.score
FROM psql_bm25s_query(
    'docs_bm25_idx'::regclass,
    'apple fruit',
    5
) AS h
JOIN docs AS d ON d.ctid = h.ctid
ORDER BY h.score DESC, d.id;
```

With no `WITH (...)` options, the index uses Lucene-style BM25 and IDF defaults with the `realtime` consistency policy.

### Indexed Inputs

`psql_bm25s` supports five indexed source-column types:

- `text` and `varchar` for direct indexing of ordinary scalar text columns.
- `text[]` and `varchar[]` for application-owned token streams.
- `int4[]` for applications that manage token IDs externally.

Scalar `text` and `varchar` columns are tokenized at the index boundary. Pretokenized arrays avoid scalar retokenization and are the preferred shape when the application already owns tokenization.

### Retrieval APIs

The canonical exact BM25 APIs are:

- `psql_bm25s_query_tokens(regclass, text[], k, weight_mask)` for token-text indexes.
- `psql_bm25s_query_ids(regclass, int4[], k, weight_mask)` for token-ID indexes.

The main SQL convenience APIs are:

- `psql_bm25s_query(regclass, query_text, k, weight_mask, ...)`
- `psql_bm25s_prepared_query(regclass, query_text, ...)`
- `psql_bm25s_query_prepared(prepared_query, k, weight_mask)`

These rowset APIs return `psql_bm25s_result_hit` rows with `ctid`, `doc_id`, and `score`. Join hits back to application rows with `ctid` when the query needs both row data and the query-time score.

### Operators

The operator surface is useful for SQL-native filtering and ordering:

- `tokens @@ 'query text'` is a boolean document-match predicate.
- `tokens @@@ psql_bm25s_prepared_query(...)` is an index-bound prepared predicate.
- `ORDER BY tokens <=> psql_bm25s_order_tokens(...) ASC LIMIT k` is the ordered retrieval form.

`@@` is not a ranking API. `<=>` aligns with true BM25 ordering only when PostgreSQL chooses a real `psql_bm25s` index scan; use the rowset retrieval APIs when you need the clearest exact top-k contract.

```sql
WITH rq AS (
    SELECT psql_bm25s_ranked_query(
        'docs_bm25_idx'::regclass,
        'apple fruit',
        10
    ) AS q
)
SELECT d.*
FROM docs AS d, rq
WHERE d.body @@@ psql_bm25s_filter_query(rq.q)
ORDER BY d.body <=> psql_bm25s_order_tokens(rq.q) ASC
LIMIT (rq.q).k;
```

### Index Options

Scoring reloptions include:

- `method` and `idf_method`, defaulting to `lucene`; supported variants are `robertson`, `lucene`, `atire`, `bm25l`, and `bm25+`.
- `k1`, default `1.5`.
- `b`, default `0.75`.
- `delta`, default `0.5`, used by BM25L and BM25+.

Scalar text processing reloptions include:

- `text_lowercase`, default `true`.
- `text_stopwords`, default `NULL`.
- `text_stem_english`, default `false`.
- `text_fold_diacritics`, default `false`.

```sql
CREATE INDEX docs_body_bm25_idx
    ON docs USING psql_bm25s (body)
    WITH (
        method = 'lucene',
        idf_method = 'lucene',
        k1 = 1.5,
        b = 0.75,
        text_stem_english = true
    );
```

### Multi-Field and Hybrid Search

For separate title, abstract, and body indexes, use late fusion helpers so each field produces query-scoped hits before scores are combined:

```sql
SELECT d.id, d.title, h.score
FROM psql_bm25s_fusion_query(
    ARRAY[
        'docs_title_bm25_idx'::regclass,
        'docs_body_bm25_idx'::regclass
    ],
    'apple fruit',
    ARRAY[3.0, 1.0]::real[],
    10,
    30,
    NULL
) AS h
JOIN docs AS d ON d.ctid = h.ctid
ORDER BY h.score DESC, d.id;
```

Hybrid vector/BM25 search is also a late-fusion layer. BM25 and vector indexes produce candidates independently, and `psql_bm25s_hybrid_fuse_candidates(...)` combines them using `rrf` by default. The core extension does not require `pgvector`, VectorChord, or any vector type to be installed.

### Maintenance and Cache

The public maintenance switch is the `consistency` reloption:

- `realtime`, the default, keeps committed writes searchable immediately.
- `eventual` favors foreground read/write latency and allows short-term stale BM25 results while maintenance converges.
- `manual` leaves refresh under explicit operator or scheduler control.

Operational helpers include:

- `psql_bm25s_index_details(regclass)`
- `psql_bm25s_index_policy_recommend(regclass, profile)`
- `psql_bm25s_index_refresh(regclass)`
- `psql_bm25s_index_maintain(regclass)`
- `psql_bm25s_index_try_maintain(regclass)`
- `psql_bm25s_index_maintain_due(max_indexes)`

Large immutable index payloads can use a shared generation cache. The zero-configuration DSM cache does not require `shared_preload_libraries`. For large connection-pool deployments, an optional shared-preload arena can be enabled with `shared_preload_libraries = 'psql_bm25s'` and `psql_bm25s.shared_generation_cache_size`, but that arena is not required for normal use.

Related global GUCs include:

- `psql_bm25s.maintenance_worker_limit`
- `psql_bm25s.preload_timer_interval_ms`
- `psql_bm25s.maintenance_timer_interval_ms`
- `psql_bm25s.maintenance_rebuild_memory_budget`

### Caveats

- The extension is not relocatable after creation; choose a non-`public` schema at `CREATE EXTENSION` time if needed.
- `eventual` and `manual` consistency deliberately trade immediate freshness for lower foreground cost or explicit refresh control.
- Logical replication follows PostgreSQL behavior: table rows replicate, but index relations do not replicate as logical data objects, so indexes should be created or rebuilt on subscribers.
- The optional shared-preload cache requires PostgreSQL configuration and a restart because the shared arena is allocated at server start.

### Upgrade to 0.4.14

After installing matching extension files, run the extension update in each database. The versioned migration from 0.4.11 to 0.4.14 declares an unchanged SQL surface; the release does not document a mandatory index rebuild.

```sql
ALTER EXTENSION psql_bm25s UPDATE TO '0.4.14';
```
