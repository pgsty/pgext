---
title: "pg_jieba"
linkTitle: "pg_jieba"
description: "Chinese full-text search parser based on cppjieba"
weight: 2240
categories: ["FTS"]
width: full
---

[**pg_jieba**](https://github.com/jaiminpan/pg_jieba) : Chinese full-text search parser based on cppjieba


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2240** | {{< badge content="pg_jieba" link="https://github.com/jaiminpan/pg_jieba" >}} | {{< ext "pg_jieba" >}} | `1.1.0` | {{< category "FTS" >}} | {{< license "BSD-3-Clause" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cjk_parser" >}} {{< ext "zhparser" >}} {{< ext "pg_bigm" >}} {{< ext "pgroonga" >}} {{< ext "pg_tokenizer" >}} |

> [!Note] Package 2.0.1 ships extension version 1.1.0, vendors cppjieba commit 45809955, and fixes the LexDescr terminator allocation.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_jieba` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "pg_jieba_18" "green" >}} {{< bg "17" "pg_jieba_17" "green" >}} {{< bg "16" "pg_jieba_16" "green" >}} {{< bg "15" "pg_jieba_15" "green" >}} {{< bg "14" "pg_jieba_14" "green" >}} | `pg_jieba_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "postgresql-18-pg-jieba" "green" >}} {{< bg "17" "postgresql-17-pg-jieba" "green" >}} {{< bg "16" "postgresql-16-pg-jieba" "green" >}} {{< bg "15" "postgresql-15-pg-jieba" "green" >}} {{< bg "14" "postgresql-14-pg-jieba" "green" >}} | `postgresql-$v-pg-jieba` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "pg_jieba_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-18-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-17-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-16-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-15-pg-jieba : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.1" "postgresql-14-pg-jieba : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jieba_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.7 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jieba_18-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_jieba_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.7 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jieba_18-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_jieba_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jieba_18-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_jieba_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jieba_18-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_jieba_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_jieba_18-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_jieba_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_jieba_18-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_jieba_18-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-jieba` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-jieba` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-18-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jieba_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.7 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jieba_17-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_jieba_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.7 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jieba_17-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_jieba_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jieba_17-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_jieba_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jieba_17-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_jieba_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_jieba_17-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_jieba_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_jieba_17-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_jieba_17-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-jieba` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-jieba` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-17-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jieba_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.7 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jieba_16-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_jieba_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.7 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jieba_16-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_jieba_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jieba_16-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_jieba_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jieba_16-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_jieba_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_jieba_16-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_jieba_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_jieba_16-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_jieba_16-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-jieba` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-jieba` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-16-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jieba_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.7 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jieba_15-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_jieba_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.7 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jieba_15-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_jieba_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jieba_15-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_jieba_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jieba_15-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_jieba_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_jieba_15-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_jieba_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_jieba_15-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_jieba_15-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-jieba` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-jieba` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-15-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_jieba_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.7 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jieba_14-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_jieba_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.7 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jieba_14-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_jieba_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jieba_14-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_jieba_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jieba_14-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_jieba_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_jieba_14-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_jieba_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_jieba_14-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_jieba_14-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-jieba` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-jieba` | `2.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-jieba/postgresql-14-pg-jieba_2.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jaiminpan/pg_jieba" title="Repository" icon="github" subtitle="github.com/jaiminpan/pg_jieba" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_jieba-2.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_jieba;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_jieba;		# install via package name, for the active PG version

pig install pg_jieba -v 18;   # install for PG 18
pig install pg_jieba -v 17;   # install for PG 17
pig install pg_jieba -v 16;   # install for PG 16
pig install pg_jieba -v 15;   # install for PG 15
pig install pg_jieba -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_jieba;
```

## Usage

Sources:

- [Official v2.0.1 README](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/README.md)
- [Extension control file](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.control)
- [SQL parser and configuration definitions](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.sql)

`pg_jieba` adds Jieba-based Chinese word segmentation to PostgreSQL full-text search. The upstream `v2.0.1` source release installs SQL extension version `1.1.0`, as recorded by its control file. It provides separate document and query parsers plus ready-to-use text-search configurations.

### Core Workflow

```sql
CREATE EXTENSION pg_jieba;

SELECT to_tsvector(
    'jiebacfg',
    '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);

SELECT plainto_tsquery('jiebaqry', '云计算专家');
```

Use `jiebacfg` to build searchable document vectors and `jiebaqry` to segment user queries:

```sql
ALTER TABLE articles
ADD COLUMN search_vector tsvector
GENERATED ALWAYS AS (to_tsvector('jiebacfg', body)) STORED;

CREATE INDEX articles_search_idx
ON articles USING GIN (search_vector);

SELECT title
FROM articles
WHERE search_vector @@ plainto_tsquery('jiebaqry', '中文全文检索');
```

### Object Index

- `jieba`: document text-search parser.
- `jiebaqry`: query-oriented text-search parser.
- `jiebacfg`: document text-search configuration using `jieba` and `jieba_stem`.
- `jiebaqry`: text-search configuration of the same name using the query parser.
- `jieba_stem`: simple dictionary with Jieba stop words used for the parser's token categories.

### Custom Dictionary and Caveats

Upstream reads a custom dictionary named `jieba.user.dict.utf8` from PostgreSQL's `tsearch_data` directory. Entries may contain a word and optional part-of-speech tag:

```text
云计算
韩玉鉴赏
蓝翔 nz
```

- The v2.x source requires a C++11-capable compiler because of its bundled `cppjieba` dependency.
- Upstream's published compatibility testing is old and limited. Build and regression-test the package against the exact PostgreSQL major version used in production.
- Changing dictionaries changes tokenization. Recompute stored `tsvector` values and rebuild dependent indexes when dictionary output changes.
