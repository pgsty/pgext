---
title: "pg_bestmatch"
linkTitle: "pg_bestmatch"
description: "Generate BM25 sparse vector inside PostgreSQL"
weight: 2140
categories: ["FTS"]
width: full
---

[**pg_bestmatch**](https://github.com/tensorchord/pg_bestmatch.rs)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2140** | {{< badge content="pg_bestmatch" link="https://github.com/tensorchord/pg_bestmatch.rs" >}} | {{< ext "pg_bestmatch" >}} | `0.0.2` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "pg_search" >}} {{< ext "vchord_bm25" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "zhparser" >}} {{< ext "pg_tokenizer" >}} {{< ext "vectorize" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_bestmatch" >}} | `0.0.2` | {{< bg "18" "pg_bestmatch_18" "green" >}} {{< bg "17" "pg_bestmatch_17" "green" >}} {{< bg "16" "pg_bestmatch_16" "green" >}} {{< bg "15" "pg_bestmatch_15" "green" >}} {{< bg "14" "pg_bestmatch_14" "green" >}} {{< bg "13" "pg_bestmatch_13" "green" >}} | `pg_bestmatch_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_bestmatch" >}} | `0.0.2` | {{< bg "18" "postgresql-18-pg-bestmatch" "green" >}} {{< bg "17" "postgresql-17-pg-bestmatch" "green" >}} {{< bg "16" "postgresql-16-pg-bestmatch" "green" >}} {{< bg "15" "postgresql-15-pg-bestmatch" "green" >}} {{< bg "14" "postgresql-14-pg-bestmatch" "green" >}} {{< bg "13" "postgresql-13-pg-bestmatch" "green" >}} | `postgresql-$v-pg-bestmatch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_bestmatch_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-bestmatch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-bestmatch : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_18` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_18` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_18` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_18` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_18` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_18-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_18` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_18-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_18-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.8 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.8 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.5 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-18-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_17` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.7 MiB | [pg_bestmatch_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_17` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_bestmatch_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_17` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_bestmatch_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_17` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_17` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_17-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_17` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_17-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_17-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.8 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.9 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-17-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_16` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.7 MiB | [pg_bestmatch_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_16` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_bestmatch_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_16` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_bestmatch_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_16` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_16` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_16-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_16` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_16-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_16-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.8 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.9 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-16-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_15` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.7 MiB | [pg_bestmatch_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_15` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_bestmatch_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_15` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_bestmatch_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_15` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_15` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_15-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_15` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_15-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_15-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.9 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.8 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.5 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-15-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_14` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.7 MiB | [pg_bestmatch_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_14` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_bestmatch_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_14` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_bestmatch_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_14` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_14` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_14-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_14` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_14-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_14-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.9 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.8 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-14-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bestmatch_13` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_13-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.7 MiB | [pg_bestmatch_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bestmatch_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_bestmatch_13` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_13-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.5 MiB | [pg_bestmatch_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bestmatch_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_bestmatch_13` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_13-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.4 MiB | [pg_bestmatch_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bestmatch_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_bestmatch_13` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bestmatch_13-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_bestmatch_13` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.9 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bestmatch_13-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_bestmatch_13` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pg_bestmatch_13-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bestmatch_13-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.1 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.8 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.1 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.8 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.6 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-bestmatch` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.5 MiB | [postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bestmatch/postgresql-13-pg-bestmatch_0.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/pg_bestmatch.rs" title="Repository" icon="github" subtitle="github.com/tensorchord/pg_bestmatch.rs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bestmatch-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_bestmatch; # get pg_bestmatch source code
pig build dep pg_bestmatch; # install build dependencies
pig build pkg pg_bestmatch; # build extension rpm or deb
pig build ext pg_bestmatch; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_bestmatch; # install by extension name, for the current active PG version
pig ext install pg_bestmatch; # install via package alias, for the active PG version
pig ext install pg_bestmatch -v 18;   # install for PG 18
pig ext install pg_bestmatch -v 17;   # install for PG 17
pig ext install pg_bestmatch -v 16;   # install for PG 16
pig ext install pg_bestmatch -v 15;   # install for PG 15
pig ext install pg_bestmatch -v 14;   # install for PG 14
pig ext install pg_bestmatch -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_bestmatch CASCADE SCHEMA bm_catalog;
```



--------

## Usage

- repo: https://github.com/tensorchord/pg_bestmatch.rs
- benchmark: https://hazyresearch.stanford.edu/blog/2024-05-20-m2-bert-retrieval


### How does it work?

- Create an BM25 statistics based on your document set by `bm25_create(table_name, column_name, statistic_name);`. It will create a materilized view to record the stats.
- Generate document sparse vector by `bm25_document_to_svector(statistic_name, passage)`
- For query, generate query sparse vector `bm25_query_to_svector(statistic_name, query)`
- Calculate the score by dot product between the query sparse vector and the document sparse vector
- Currently we use huggingface tokenizer with `bert-base-uncased` vocabulary set to tokenize words. Might support more configuration on tokenizer in the future.


### Install

```sql
CREATE EXTENSION pg_bestmatch;
SET search_path TO public, bm_catalog;
```



--------

## Example

Here is an example workflow demonstrating the usage of this extension with the example of [Stanford LoCo benchmark](https://hazyresearch.stanford.edu/blog/2024-05-20-m2-bert-retrieval).

0. Load the dataset. Here is a script for you if you want to experience `pg_bestmatch` with the dataset.

```sh
wget https://huggingface.co/api/datasets/hazyresearch/LoCoV1-Documents/parquet/default/test/0.parquet -O documents.parquet
wget https://huggingface.co/api/datasets/hazyresearch/LoCoV1-Queries/parquet/default/test/0.parquet -O queries.parquet
```

```python
import pandas as pd
from sqlalchemy import create_engine
import numpy as np
from psycopg2.extensions import register_adapter, AsIs

def adapter_numpy_float64(numpy_float64):
    return AsIs(numpy_float64)

def adapter_numpy_int64(numpy_int64):
    return AsIs(numpy_int64)

def adapter_numpy_float32(numpy_float32):
    return AsIs(numpy_float32)

def adapter_numpy_int32(numpy_int32):
    return AsIs(numpy_int32)

def adapter_numpy_array(numpy_array):
    return AsIs(tuple(numpy_array))

register_adapter(np.float64, adapter_numpy_float64)
register_adapter(np.int64, adapter_numpy_int64)
register_adapter(np.float32, adapter_numpy_float32)
register_adapter(np.int32, adapter_numpy_int32)
register_adapter(np.ndarray, adapter_numpy_array)

db_url = "postgresql://localhost:5432/pg_bestmatch_test"
engine = create_engine(db_url)

def load_documents():
    df = pd.read_parquet("documents.parquet")
    df.to_sql("documents", engine, if_exists='replace', index=False)

def load_queries():
    df = pd.read_parquet("queries.parquet")
    df['answer_pids'] = df['answer_pids'].apply(lambda x: str(x[0]))    
    df.to_sql("queries", engine, if_exists='replace', index=False)

load_documents()
load_queries()
```

1. Create BM25 statistics for the `documents` table.

```sql
SELECT bm25_create('documents', 'passage', 'documents_passage_bm25', 0.75, 1.2);
```

2. Add an embedding column to the `documents` and `queries` tables and update the embeddings for documents and queries.

```sql
ALTER TABLE documents ADD COLUMN embedding svector; -- for pgvecto.rs users
ALTER TABLE documents ADD COLUMN embedding sparsevec; -- for pgvector users

UPDATE documents SET embedding = bm25_document_to_svector('documents_passage_bm25', passage)::svector; -- for pgvecto.rs users
UPDATE documents SET embedding = bm25_document_to_svector('documents_passage_bm25', passage, 'pgvector')::sparsevec; -- for pgvector users
```

3. (Optional) Create a vector index on the sparse vector column.

```sql
CREATE INDEX ON documents USING vectors (embedding svector_dot_ops); -- for pgvecto.rs users
CREATE INDEX ON documents USING ivfflat (embedding sparsevec_ip_ops); -- for pgvector users
```

4. Perform a vector search to find the most relevant documents for each query.

```sql
ALTER TABLE queries ADD COLUMN embedding svector; -- for pgvecto.rs users
ALTER TABLE queries ADD COLUMN embedding sparsevec; -- for pgvector users

UPDATE queries SET embedding = bm25_query_to_svector('documents_passage_bm25', query)::svector; -- for pgvecto.rs users
UPDATE queries SET embedding = bm25_query_to_svector('documents_passage_bm25', query, 'pgvector')::sparsevec; -- for pgvector users

SELECT sum((array[answer_pids] = array(SELECT pid FROM documents WHERE queries.dataset = documents.dataset ORDER BY queries.embedding <#> documents.embedding LIMIT 1))::int) FROM queries;
```

This workflow showcases how to leverage BM25 text queries and vector search in PostgreSQL using this extension. The Top 1 recall of BM25 on this dataset is `0.77`. If you reproduce the result, your operations are correct.


--------

## Comparison with pg_search

- `pg_bestmatch.rs` only provides methods for generating sparse vectors and does not support index-based search (which can be achieved by pgvecto.rs or pgvector).
- `pg_search` performs BM25 retrieval via the external `tantivy` engine, which may have limitations when combined with transactions, filters, or JOIN operations. Since `pg_bestmatch.rs` is entirely native to Postgres, it offers full compatibility with these operations inside postgres.


--------

## Reference

- `tokenize`
    - Description: Tokenizes an input string into individual tokens.
    - Example:
      ```sql
      SELECT tokenize('i have an apple'); -- result: {i,have,an,apple}
      ```
- `bm25_create`
    - Description: Creates BM25 statistics for a specified table and column.
    - Usage:
      ```sql
      SELECT bm25_create('documents', 'passage', 'documents_passage_bm25');
      ```
    - Parameters:
        - `table_name`: Name of the table.
        - `column_name`: Name of the column.
        - `stat_name`: Name of the BM25 statistics.
        - `b`: BM25 parameter (default 0.75).
        - `k`: BM25 parameter (default 1.2).
- `bm25_refresh`
    - Description: Updates the BM25 statistics to reflect any changes in the underlying data.
    - Usage:
      ```sql
      SELECT bm25_refresh('documents_passage_bm25');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics to update.
- `bm25_drop`
    - Description: Deletes the BM25 statistics for a specified table and column.
    - Usage:
      ```sql
      SELECT bm25_drop('documents_passage_bm25');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics to delete.
- `bm25_document_to_svector`
    - Description: Converts document text into a sparse vector representation.
    - Usage:
      ```sql
      SELECT bm25_document_to_svector('documents_passage_bm25', 'document_text');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics.
        - `document_text`: The text of the document.
        - `style`: Emits `pgvecto.rs`-style sparse vector or `pgvector`-style sparse vector.
- `bm25_query_to_svector`
    - Description: Converts query text into a sparse vector representation.
    - Usage:
      ```sql
      SELECT bm25_query_to_svector('documents_passage_bm25', 'We begin, as always, with the text.');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics.
        - `query_text`: The text of the query.
        - `style`: Emits `pgvecto.rs`-style sparse vector or `pgvector`-style sparse vector.