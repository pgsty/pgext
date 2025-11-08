---
title: "pg_tokenizer"
linkTitle: "pg_tokenizer"
description: "Tokenizers for full-text search"
weight: 2160
categories: ["FTS"]
width: full
---

Tokenizers for full-text search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2160** | {{< badge content="pg_tokenizer" link="https://github.com/tensorchord/pg_tokenizer.rs" >}} | {{< ext "pg_tokenizer" >}} | `0.1.1` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_trgm" >}} |

> [!Note] PG18 fix by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_tokenizer" >}} | `0.1.1` | {{< bg "18" "pg_tokenizer_18" "green" >}} {{< bg "17" "pg_tokenizer_17" "green" >}} {{< bg "16" "pg_tokenizer_16" "green" >}} {{< bg "15" "pg_tokenizer_15" "green" >}} {{< bg "14" "pg_tokenizer_14" "green" >}} {{< bg "13" "pg_tokenizer_13" "green" >}} | `pg_tokenizer_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tokenizer" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pg-tokenizer" "green" >}} {{< bg "17" "postgresql-17-pg-tokenizer" "green" >}} {{< bg "16" "postgresql-16-pg-tokenizer" "green" >}} {{< bg "15" "postgresql-15-pg-tokenizer" "green" >}} {{< bg "14" "postgresql-14-pg-tokenizer" "green" >}} {{< bg "13" "postgresql-13-pg-tokenizer" "green" >}} | `postgresql-$v-pg-tokenizer` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_tokenizer_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pg-tokenizer : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-13-pg-tokenizer : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_18` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_18` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_18` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_18` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_18` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_18` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 5.0 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 5.0 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 5.0 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 5.0 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-18-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_17` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_17` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_17` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_17` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_17` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_17` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 9.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 9.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 9.9 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 9.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 10.8 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 10.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 10.8 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 10.7 MiB | [postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-17-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_16` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_16` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_16` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_16` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_16` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_16` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 10.0 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 9.8 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 9.9 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 9.8 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 10.8 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 10.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 10.8 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 10.7 MiB | [postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-16-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_15` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_15` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_15` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_15` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_15` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_15-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_15` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_15-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_15-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 9.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 9.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 9.9 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 9.8 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 10.8 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 10.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 10.8 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 10.7 MiB | [postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-15-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_14` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_14` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_14` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_14` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_14` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_14-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_14` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_14-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_14-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 9.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 9.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 9.9 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 9.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 10.8 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 10.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 10.8 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 10.7 MiB | [postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-14-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tokenizer_13` | 0.1.1 | `el8.x86_64` | pigsty | 11.7 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tokenizer_13-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tokenizer_13` | 0.1.1 | `el8.aarch64` | pigsty | 11.5 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tokenizer_13-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tokenizer_13` | 0.1.1 | `el9.x86_64` | pigsty | 11.0 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tokenizer_13-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tokenizer_13` | 0.1.1 | `el9.aarch64` | pigsty | 10.9 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tokenizer_13-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tokenizer_13` | 0.1.1 | `el10.x86_64` | pigsty | 10.9 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tokenizer_13-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tokenizer_13` | 0.1.1 | `el10.aarch64` | pigsty | 11.0 MiB | [pg_tokenizer_13-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tokenizer_13-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `d12.x86_64` | pigsty | 5.0 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `d12.aarch64` | pigsty | 5.0 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `d13.x86_64` | pigsty | 5.0 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `d13.aarch64` | pigsty | 5.0 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-tokenizer` | 0.1.1 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tokenizer/postgresql-13-pg-tokenizer_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/pg_tokenizer.rs" title="Repository" icon="github" subtitle="github.com/tensorchord/pg_tokenizer.rs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tokenizer.rs-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_tokenizer; # get pg_tokenizer source code
pig build dep pg_tokenizer; # install build dependencies
pig build pkg pg_tokenizer; # build extension rpm or deb
pig build ext pg_tokenizer; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_tokenizer; # install by extension name, for the current active PG version
pig ext install pg_tokenizer; # install via package alias, for the active PG version
pig ext install pg_tokenizer -v 18;   # install for PG 18
pig ext install pg_tokenizer -v 17;   # install for PG 17
pig ext install pg_tokenizer -v 16;   # install for PG 16
pig ext install pg_tokenizer -v 15;   # install for PG 15
pig ext install pg_tokenizer -v 14;   # install for PG 14
pig ext install pg_tokenizer -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_tokenizer CASCADE SCHEMA tokenizer_catalog;
```

