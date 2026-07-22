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
| **2190** | {{< badge content="pg_pinyin" link="https://github.com/aiyou178/pg_pinyin" >}} | {{< ext "pg_pinyin" >}} | `0.0.5` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pinyin` |
|   **See Also**    | {{< ext "zhparser" >}} {{< ext "pg_search" >}} {{< ext "pg_trgm" >}} {{< ext "pg_bigm" >}} {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_tokenizer" >}} {{< ext "fuzzystrmatch" >}} |

> [!Note] optional tokenizer-input overload can integrate with pg_search.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_pinyin` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "pg_pinyin_18" "green" >}} {{< bg "17" "pg_pinyin_17" "green" >}} {{< bg "16" "pg_pinyin_16" "green" >}} {{< bg "15" "pg_pinyin_15" "green" >}} {{< bg "14" "pg_pinyin_14" "green" >}} | `pg_pinyin_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.5` | {{< bg "18" "postgresql-18-pinyin" "green" >}} {{< bg "17" "postgresql-17-pinyin" "green" >}} {{< bg "16" "postgresql-16-pinyin" "green" >}} {{< bg "15" "postgresql-15-pinyin" "green" >}} {{< bg "14" "postgresql-14-pinyin" "green" >}} | `postgresql-$v-pinyin` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "pg_pinyin_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-18-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-17-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-16-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-15-pinyin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.5" "postgresql-14-pinyin : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_18` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_18-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_18-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_18` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_18-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_18-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_18` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_18-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_18` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_18-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_18-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pinyin` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.5 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.5 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pinyin` | `0.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-18-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_17` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_17-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_17-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_17` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_17-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_17-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_17` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_17-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_17` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_17-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_17-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pinyin` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.5 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.5 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pinyin` | `0.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-17-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_16` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_16-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_16-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_16` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_16-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_16-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_16` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_16-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_16` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_16-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_16-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pinyin` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.5 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.5 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pinyin` | `0.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-16-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-16-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_15` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_15-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_15-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_15` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_15-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_15-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_15` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_15-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_15` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_15-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_15-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pinyin` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.5 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.5 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pinyin` | `0.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.5 MiB | [postgresql-15-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-15-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pinyin_14` | `0.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pinyin_14-0.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pinyin_14-0.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_pinyin_14` | `0.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pinyin_14-0.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pinyin_14-0.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_pinyin_14` | `0.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pinyin_14-0.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pg_pinyin_14` | `0.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [pg_pinyin_14-0.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pinyin_14-0.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pinyin` | `0.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.5 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.5 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pinyin` | `0.0.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.5 MiB | [postgresql-14-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pinyin/postgresql-14-pinyin_0.0.5-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aiyou178/pg_pinyin" title="Repository" icon="github" subtitle="github.com/aiyou178/pg_pinyin" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_pinyin-0.0.5.tar.gz" >}}
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

Sources:

- [pg_pinyin v0.0.5 README](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/readme.md)
- [pg_pinyin v0.0.5 control file](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin.control)
- [0.0.4 to 0.0.5 upgrade SQL](https://github.com/aiyou178/pg_pinyin/blob/v0.0.5/pg_pinyin--0.0.4--0.0.5.sql)

pg_pinyin romanizes Chinese text and exposes tokenizer and query helpers for search applications. Use pg_pinyin to create stable Pinyin search keys, tokenize Han text, or expand Pinyin input into a pg_search regular-expression query.

Version 0.0.5 is primarily a packaging and toolchain update; its upgrade script makes no SQL catalog changes, so the user-facing API remains compatible with 0.0.4.

### Create the Extension

    CREATE EXTENSION pg_pinyin;

The extension is relocatable and does not require shared_preload_libraries or a server restart.

### Romanize Text

Romanize character by character or use word-aware segmentation:

    SELECT pinyin_char_romanize('重庆');
    SELECT pinyin_word_romanize('重庆火锅');
    SELECT pinyin_word_romanize('重庆火锅', ' ');

Both functions accept an optional suffix inserted after each emitted Pinyin unit. Character mode is deterministic per character; word mode uses the bundled word dictionary to resolve contextual pronunciations.

### Use pg_search Tokenizer Input

Word romanization also accepts a pg_search tokenizer result when that extension is available:

    SELECT pinyin_word_romanize(
      description::pdb.icu::text[]
    )
    FROM documents;

The overload returns romanized text; it does not expose a row-per-token API. Use the plain-text overload when pg_search tokenization is not required.

### Build a pg_search Query

When pg_search was installed before pg_pinyin, pg_pinyin provides a typed overload that returns pdb.query:

    SELECT *
    FROM documents
    WHERE id @@@ pinyin_regex_phrase(
      'chong qing',
      slope => 1,
      max_expansions => 64,
      generated_pinyin => true
    );

If pg_search is absent, the same entry point is installed as an error-reporting stub rather than silently returning a different type. Install dependencies in the intended order and test the function signature after upgrades.

### Object Index

- pinyin_char_romanize(text [, suffix]) returns character-based Pinyin text.
- pinyin_word_romanize(text [, suffix]) returns dictionary-segmented Pinyin text.
- pinyin_word_romanize(tokenizer_input [, suffix]) accepts a pg_search tokenizer result.
- pinyin_regex_phrase(text, slope, max_expansions, generated_pinyin) constructs a pg_search Pinyin phrase query when that integration is available.
- pinyin_regex_phrase_patterns is an internal pattern-building helper; prefer the public query function.

### Operational Notes

The extension ships generated character and word dictionaries in its pinyin schema. Treat those tables as extension-managed data rather than application tables. Romanization is normalization, not translation, and ambiguous or domain-specific readings may require application-side review.
