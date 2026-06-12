---
title: "pg_pinyin"
linkTitle: "pg_pinyin"
description: "Pinyin romanization and search helpers for PostgreSQL"
weight: 2190
categories: ["FTS"]
width: full
---

[**pg_pinyin**](https://github.com/aiyou178/pg_pinyin) : Pinyin romanization and search helpers for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2190** | {{< badge content="pg_pinyin" link="https://github.com/aiyou178/pg_pinyin" >}} | {{< ext "pg_pinyin" >}} | `0.0.4` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pinyin` |
|   **See Also**    | {{< ext "zhparser" >}} {{< ext "pg_search" >}} {{< ext "pg_trgm" >}} {{< ext "pg_bigm" >}} {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_tokenizer" >}} {{< ext "fuzzystrmatch" >}} |

> [!Note] pgrx 0.17.0; optional tokenizer-input overload can integrate with pg_search


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_pinyin` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "pg_pinyin_18" "green" >}} {{< bg "17" "pg_pinyin_17" "green" >}} {{< bg "16" "pg_pinyin_16" "green" >}} {{< bg "15" "pg_pinyin_15" "green" >}} {{< bg "14" "pg_pinyin_14" "green" >}} | `pg_pinyin_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "postgresql-18-pinyin" "green" >}} {{< bg "17" "postgresql-17-pinyin" "green" >}} {{< bg "16" "postgresql-16-pinyin" "green" >}} {{< bg "15" "postgresql-15-pinyin" "green" >}} {{< bg "14" "postgresql-14-pinyin" "green" >}} | `postgresql-$v-pinyin` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pinyin : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pinyin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pinyin : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_18` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_18` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_18` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_18-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_18-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_18-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pinyin` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.3 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pinyin_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pinyin_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_17` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_17` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_17` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_17-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_17-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_17-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pinyin` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.3 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pinyin_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pinyin_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_16` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_16` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_16` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_16-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_16-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_16-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pinyin` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.3 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pinyin_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pinyin_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_15` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_15` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_15` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_15-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_15-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_15-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pinyin` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.3 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pinyin_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-15-pinyin_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_14` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_14` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_14` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_14-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_pinyin_14-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_14-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pinyin` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.3 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pinyin_0.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-14-pinyin_0.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aiyou178/pg_pinyin" title="Repository" icon="github" subtitle="github.com/aiyou178/pg_pinyin" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_pinyin-0.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_pinyin;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_pinyin;		# install via package name, for the active PG version

pig install pg_pinyin -v 18;   # install for PG 18
pig install pg_pinyin -v 17;   # install for PG 17
pig install pg_pinyin -v 16;   # install for PG 16
pig install pg_pinyin -v 15;   # install for PG 15
pig install pg_pinyin -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_pinyin;
```

## Usage

> Sources: [pg_pinyin upstream README](https://github.com/aiyou178/pg_pinyin), [Chinese README](https://github.com/aiyou178/pg_pinyin/blob/main/README.zh-CN.md), [local metadata](../db/extension.csv).

`pg_pinyin` converts Chinese text to Pinyin, either character by character or by word. It is useful for generated search columns, trigram search, and `pg_search` BM25 queries that need Pinyin input.

```sql
CREATE EXTENSION pg_pinyin;
```

### Functions

| Function | Description |
|----------|-------------|
| `pinyin_char_romanize(text)` | Character-level Pinyin romanization |
| `pinyin_char_romanize(text, suffix text)` | Character-level romanization with a custom dictionary suffix |
| `pinyin_word_romanize(text)` | Word-level Pinyin romanization |
| `pinyin_word_romanize(text, suffix text)` | Word-level romanization with a custom dictionary suffix |
| `pinyin_word_romanize(tokenizer_input anyelement)` | Word-level romanization from a `pg_search` tokenizer input such as `name::pdb.icu::text[]` |
| `pinyin_word_romanize(tokenizer_input anyelement, suffix text)` | Tokenizer-input romanization with a custom dictionary suffix |
| `pinyin_regex_phrase(text, slope integer DEFAULT NULL, max_expansions integer DEFAULT NULL, generated_pinyin boolean DEFAULT false)` | `pg_search` query helper returning `pdb.query`, available when `pg_search` was enabled before `CREATE EXTENSION pg_pinyin` |
| `pinyin_regex_phrase_patterns(text, generated_pinyin boolean DEFAULT false)` | Internal helper returning regex phrase tokens as `text[]` |

### Generated Column + Trigram Search

```sql
CREATE EXTENSION IF NOT EXISTS pg_pinyin;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_char_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_trgm_idx ON voice USING gin (pinyin gin_trgm_ops);

INSERT INTO voice (description) VALUES ('郑爽ABC');
SELECT id, description, pinyin FROM voice;
```

### Word Tokenization + pg_search

For word-oriented search, use `pinyin_word_romanize`. When `pg_search` is available, it can consume tokenizer input such as `pdb.icu::text[]`.

```sql
CREATE EXTENSION IF NOT EXISTS pg_search;
CREATE EXTENSION IF NOT EXISTS pg_pinyin;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_word_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_bm25_idx
ON voice
USING bm25 (id, pinyin)
WITH (key_field='id');

SELECT *
FROM voice
WHERE pinyin @@@ public.pinyin_regex_phrase('zhengshuang');

SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[]);
```

`pinyin_regex_phrase` has return type `pdb.query`, so `pg_search` must be enabled in the database before `pg_pinyin` is created. If `pg_pinyin` is created first, upstream documents that the romanization functions are installed, but `pinyin_regex_phrase` is installed as an error stub with a clear exception.

### Dictionary Tables

The extension seeds bundled dictionary tables under schema `pinyin` during `CREATE EXTENSION pg_pinyin`; no separate data-load step is needed for normal extension usage. The bundled data covers character mappings, word tokens, and word mappings.

Provide custom dictionary tables in schema `pinyin` with a suffix. Calls using that suffix merge the base dictionary with the suffix tables, and suffix entries take priority.

```sql
CREATE TABLE IF NOT EXISTS pinyin.pinyin_mapping_suffix1 (
  character text PRIMARY KEY,
  pinyin text NOT NULL
);

CREATE TABLE IF NOT EXISTS pinyin.pinyin_words_suffix1 (
  word text PRIMARY KEY,
  pinyin text NOT NULL
);

INSERT INTO pinyin.pinyin_mapping_suffix1 (character, pinyin)
VALUES ('郑', '|zhengx|')
ON CONFLICT (character) DO UPDATE SET pinyin = EXCLUDED.pinyin;

INSERT INTO pinyin.pinyin_words_suffix1 (word, pinyin)
VALUES ('郑爽', '|zhengx| |shuangx|')
ON CONFLICT (word) DO UPDATE SET pinyin = EXCLUDED.pinyin;

SELECT public.pinyin_char_romanize('郑爽ABC', '_suffix1');
SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[], '_suffix1');
```
