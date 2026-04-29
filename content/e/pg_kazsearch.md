---
title: "pg_kazsearch"
linkTitle: "pg_kazsearch"
description: "Kazakh full-text search extension for PostgreSQL"
weight: 2200
categories: ["FTS"]
width: full
---

[**pg_kazsearch**](https://github.com/darkhanakh/pg-kazsearch) : Kazakh full-text search extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2200** | {{< badge content="pg_kazsearch" link="https://github.com/darkhanakh/pg-kazsearch" >}} | {{< ext "pg_kazsearch" >}} | `0.1.0` | {{< category "FTS" >}} | {{< license "LGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Upstream release/package version is 2.0.0; extension control version is 0.1.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_kazsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "pg_kazsearch_18" "green" >}} {{< bg "17" "pg_kazsearch_17" "green" >}} {{< bg "16" "pg_kazsearch_16" "green" >}} {{< bg "15" "pg_kazsearch_15" "red" >}} {{< bg "14" "pg_kazsearch_14" "red" >}} | `pg_kazsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "postgresql-18-pg-kazsearch" "green" >}} {{< bg "17" "postgresql-17-pg-kazsearch" "green" >}} {{< bg "16" "postgresql-16-pg-kazsearch" "green" >}} {{< bg "15" "postgresql-15-pg-kazsearch" "red" >}} {{< bg "14" "postgresql-14-pg-kazsearch" "red" >}} | `postgresql-$v-pg-kazsearch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 443.6 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_18-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 321.3 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_18-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.3 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_18-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 331.2 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_18-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.1 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_18-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.3 KiB | [pg_kazsearch_18-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_18-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 361.3 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 251.2 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 361.6 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 251.3 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.8 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 285.9 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 402.2 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 284.0 KiB | [postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 443.2 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 321.4 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.4 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 331.3 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.7 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_17-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.3 KiB | [pg_kazsearch_17-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_17-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 361.5 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 251.2 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 361.5 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 251.0 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.8 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 286.3 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 401.9 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 283.9 KiB | [postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 443.4 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 321.5 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.4 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 331.4 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.2 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_16-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 331.2 KiB | [pg_kazsearch_16-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_16-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 361.5 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 251.2 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 361.2 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 251.2 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.9 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 286.2 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 402.0 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 283.6 KiB | [postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/darkhanakh/pg-kazsearch" title="Repository" icon="github" subtitle="github.com/darkhanakh/pg-kazsearch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_kazsearch-2.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_kazsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_kazsearch;		# install via package name, for the active PG version

pig install pg_kazsearch -v 18;   # install for PG 18
pig install pg_kazsearch -v 17;   # install for PG 17
pig install pg_kazsearch -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_kazsearch;
```

## Usage

Sources: [README](https://github.com/darkhanakh/pg-kazsearch/blob/main/README.md), [releases](https://github.com/darkhanakh/pg-kazsearch/releases)

`pg_kazsearch` is a PostgreSQL full-text search extension for the Kazakh language. The README says it creates a ready-to-use text search configuration `kazakh_cfg` and dictionary `pg_kazsearch_dict`.

### Quick start

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1
```

### Add Kazakh FTS to a table

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
    GENERATED ALWAYS AS (
        setweight(to_tsvector('kazakh_cfg', title), 'A') ||
        setweight(to_tsvector('kazakh_cfg', body), 'B')
    ) STORED;

CREATE INDEX idx_fts ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(fts, websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')) DESC
LIMIT 10;
```

### Tuning

The README documents runtime dictionary tuning without restart:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (w_deriv = 3.5, w_short_char = 100.0);
```

### Release and packaging notes

- Upstream release `v2.0.0` introduced the current Rust / `pgrx` architecture.
- Upstream release `v2.1.0` adds an Elasticsearch plugin alongside the PostgreSQL extension; the PostgreSQL SQL usage in the README is unchanged.
- The repository README publishes Debian packages as `2.x` releases, while this project's CSV note separately tracks the extension control version.

### Caveat

The PostgreSQL-facing docs are concise and focused on stemming plus FTS usage. For this stub, avoid inferring extra SQL objects beyond `kazakh_cfg`, `pg_kazsearch_dict`, and the documented examples above.
