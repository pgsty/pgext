---
title: "pg_cjk_parser"
linkTitle: "pg_cjk_parser"
description: "CJK bigram parser derived from PostgreSQL full-text search"
weight: 2230
categories: ["FTS"]
width: full
---

[**pg_cjk_parser**](https://github.com/huangjimmy/pg_cjk_parser) : CJK bigram parser derived from PostgreSQL full-text search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2230** | {{< badge content="pg_cjk_parser" link="https://github.com/huangjimmy/pg_cjk_parser" >}} | {{< ext "pg_cjk_parser" >}} | `0.1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jieba" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pgroonga" >}} {{< ext "pg_tokenizer" >}} |

> [!Note] PGSTY applies a PG_CONFIG build-selection patch.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_cjk_parser` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_cjk_parser_18" "green" >}} {{< bg "17" "pg_cjk_parser_17" "green" >}} {{< bg "16" "pg_cjk_parser_16" "green" >}} {{< bg "15" "pg_cjk_parser_15" "green" >}} {{< bg "14" "pg_cjk_parser_14" "green" >}} | `pg_cjk_parser_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-cjk-parser" "green" >}} {{< bg "17" "postgresql-17-pg-cjk-parser" "green" >}} {{< bg "16" "postgresql-16-pg-cjk-parser" "green" >}} {{< bg "15" "postgresql-15-pg-cjk-parser" "green" >}} {{< bg "14" "postgresql-14-pg-cjk-parser" "green" >}} | `postgresql-$v-pg-cjk-parser` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_cjk_parser_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-cjk-parser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-cjk-parser : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cjk_parser_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 36.4 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cjk_parser_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cjk_parser_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 36.0 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cjk_parser_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cjk_parser_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 37.9 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cjk_parser_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cjk_parser_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cjk_parser_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cjk_parser_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.0 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cjk_parser_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cjk_parser_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cjk_parser_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 68.1 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.6 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 68.1 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.8 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.9 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.7 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.6 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 72.5 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.7 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-cjk-parser` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.6 KiB | [postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-18-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cjk_parser_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 36.4 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cjk_parser_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cjk_parser_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 36.0 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cjk_parser_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cjk_parser_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 37.9 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cjk_parser_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cjk_parser_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cjk_parser_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cjk_parser_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.1 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cjk_parser_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cjk_parser_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cjk_parser_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 68.0 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.6 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 68.1 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.7 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.3 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 76.1 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.7 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 72.6 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.9 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-cjk-parser` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.7 KiB | [postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-17-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cjk_parser_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 36.4 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cjk_parser_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cjk_parser_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 36.0 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cjk_parser_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cjk_parser_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 37.9 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cjk_parser_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cjk_parser_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cjk_parser_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cjk_parser_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.0 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cjk_parser_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cjk_parser_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 38.1 KiB | [pg_cjk_parser_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cjk_parser_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 68.0 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.6 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 68.1 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.6 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.1 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 76.1 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.7 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 72.6 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.9 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-cjk-parser` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.7 KiB | [postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-16-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cjk_parser_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 36.2 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cjk_parser_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cjk_parser_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.9 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cjk_parser_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cjk_parser_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 38.3 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cjk_parser_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cjk_parser_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 38.0 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cjk_parser_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cjk_parser_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.3 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cjk_parser_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cjk_parser_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 38.2 KiB | [pg_cjk_parser_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cjk_parser_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.7 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.1 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.7 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.1 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.0 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.9 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.8 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 72.3 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.3 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-cjk-parser` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.3 KiB | [postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-15-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cjk_parser_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 36.2 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cjk_parser_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cjk_parser_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.9 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cjk_parser_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cjk_parser_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 38.3 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cjk_parser_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cjk_parser_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 38.0 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cjk_parser_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cjk_parser_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 38.3 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cjk_parser_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cjk_parser_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 38.4 KiB | [pg_cjk_parser_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cjk_parser_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.7 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.0 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.7 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.0 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.2 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.8 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.6 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 72.3 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.4 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-cjk-parser` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.3 KiB | [postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-cjk-parser/postgresql-14-pg-cjk-parser_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/huangjimmy/pg_cjk_parser" title="Repository" icon="github" subtitle="github.com/huangjimmy/pg_cjk_parser" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cjk_parser-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_cjk_parser;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cjk_parser;		# install via package name, for the active PG version

pig install pg_cjk_parser -v 18;   # install for PG 18
pig install pg_cjk_parser -v 17;   # install for PG 17
pig install pg_cjk_parser -v 16;   # install for PG 16
pig install pg_cjk_parser -v 15;   # install for PG 15
pig install pg_cjk_parser -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cjk_parser;
```

## Usage

Sources:

- [Official v0.1.0 README](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/Readme.md)
- [v0.1.0 release notes](https://github.com/huangjimmy/pg_cjk_parser/releases/tag/v0.1.0)
- [v0.1.0 control file](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser.control)
- [v0.1.0 SQL functions](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser--0.1.0.sql)

`pg_cjk_parser` is a PostgreSQL full-text-search parser derived from the built-in parser. In a UTF-8 database it keeps the default behavior for non-CJK text while emitting overlapping 2-gram tokens for Chinese, Japanese, and Korean text. The extension installs parser support functions; you create the text-search parser and configuration that use them.

### Core Workflow

```sql
CREATE EXTENSION pg_cjk_parser;

CREATE TEXT SEARCH PARSER public.pg_cjk_parser (
    START = prsd2_cjk_start,
    GETTOKEN = prsd2_cjk_nexttoken,
    END = prsd2_cjk_end,
    LEXTYPES = prsd2_cjk_lextype,
    HEADLINE = prsd2_cjk_headline
);

CREATE TEXT SEARCH CONFIGURATION public.config_2_gram_cjk (
    PARSER = public.pg_cjk_parser
);

SELECT alias, description, token
FROM ts_debug(
    'public.config_2_gram_cjk',
    'PostgreSQL 全文検索和中文检索'
);
```

Use the configuration explicitly in generated `tsvector` columns and queries, or set it as the session default:

```sql
SET default_text_search_config = 'public.config_2_gram_cjk';

SELECT to_tsvector('public.config_2_gram_cjk', '日本語全文検索');
```

### Important Objects

- `prsd2_cjk_start`, `prsd2_cjk_nexttoken`, `prsd2_cjk_end`, `prsd2_cjk_lextype`, and `prsd2_cjk_headline`: support functions used by `CREATE TEXT SEARCH PARSER`.
- `cjk_zht2zhs(text)`: converts mapped Traditional Chinese characters to Simplified Chinese while leaving other characters unchanged.
- Parser token type `cjk`: emits overlapping CJK bigrams; CJK punctuation is emitted as a unigram.

```sql
SELECT cjk_zht2zhs('漢語');
-- 汉语
```

### Version Notes and Caveats

- Version `0.1.0` fixes incorrect `cjk_zht2zhs` scanning across mixed-width UTF-8 characters and corrects handling of four-byte CJK code points.
- Upstream supports PostgreSQL 11 through 18 at this release.
- The database must use UTF-8 for CJK bigram behavior. With another encoding, the parser behaves like the PostgreSQL default parser.
- Creating a text-search parser requires elevated privileges. Decide mappings, dictionaries, stop words, and ranking separately; the example configuration defines only the parser.
