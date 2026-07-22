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
| **2200** | {{< badge content="pg_kazsearch" link="https://github.com/darkhanakh/pg-kazsearch" >}} | {{< ext "pg_kazsearch" >}} | `2.3.0` | {{< category "FTS" >}} | {{< license "LGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Upstream 2.3.0 uses pgrx 0.17.0; PIGSTY packaging builds with pgrx 0.19.1 for PostgreSQL 16 through 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_kazsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "pg_kazsearch_18" "green" >}} {{< bg "17" "pg_kazsearch_17" "green" >}} {{< bg "16" "pg_kazsearch_16" "green" >}} {{< bg "15" "pg_kazsearch_15" "red" >}} {{< bg "14" "pg_kazsearch_14" "red" >}} | `pg_kazsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3.0` | {{< bg "18" "postgresql-18-pg-kazsearch" "green" >}} {{< bg "17" "postgresql-17-pg-kazsearch" "green" >}} {{< bg "16" "postgresql-16-pg-kazsearch" "green" >}} {{< bg "15" "postgresql-15-pg-kazsearch" "red" >}} {{< bg "14" "postgresql-14-pg-kazsearch" "red" >}} | `postgresql-$v-pg-kazsearch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_kazsearch_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_kazsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_kazsearch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-18-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-kazsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-kazsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-kazsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-kazsearch : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_18` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_18-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_18` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1016.2 KiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_18-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_18` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_18-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_18` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_18-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_18` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_18-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_18` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_18-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_18-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 858.1 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 749.5 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 858.2 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.5 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 959.5 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 883.4 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 953.6 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 869.6 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 945.3 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-kazsearch` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 867.0 KiB | [postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-18-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_17` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_17-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_17` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1013.8 KiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_17-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_17` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_17-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_17` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_17-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_17` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_17-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_17` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_17-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_17-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 856.6 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 748.2 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 857.0 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 748.6 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 958.8 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 881.4 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 948.8 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 868.6 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 941.3 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-kazsearch` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 866.1 KiB | [postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-17-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kazsearch_16` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kazsearch_16-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_kazsearch_16` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1012.5 KiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kazsearch_16-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_kazsearch_16` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kazsearch_16-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_kazsearch_16` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kazsearch_16-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_kazsearch_16` | `2.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kazsearch_16-2.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_kazsearch_16` | `2.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_kazsearch_16-2.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kazsearch_16-2.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 856.5 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 747.3 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 853.8 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 747.9 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 956.7 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 881.4 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 954.9 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 867.5 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 941.0 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-kazsearch` | `2.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 863.5 KiB | [postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kazsearch/postgresql-16-pg-kazsearch_2.3.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/darkhanakh/pg-kazsearch" title="Repository" icon="github" subtitle="github.com/darkhanakh/pg-kazsearch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_kazsearch-2.3.0.tar.gz" >}}
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

Sources:

- [Official v2.3.0 README](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/README.md)
- [v2.3.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.3.0)
- [PostgreSQL extension control file](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/pg_kazsearch.control)
- [v2.2.0 to v2.3.0 upgrade SQL](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/sql/pg_kazsearch--2.2.0--2.3.0.sql)

`pg_kazsearch` provides Kazakh full-text stemming for PostgreSQL 16 through 18. It installs a ready-to-use `kazakh_cfg` configuration and `pg_kazsearch_dict` dictionary. Cyrillic and supported modern Latin-script Kazakh converge to canonical Cyrillic stems so documents and queries can match across scripts.

### Core Workflow

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'мектептеріміздегі оқушылардың');
-- 'мектеп':1 'оқушы':2
```

Add a weighted stored vector and GIN index:

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
GENERATED ALWAYS AS (
    setweight(to_tsvector('kazakh_cfg', title), 'A') ||
    setweight(to_tsvector('kazakh_cfg', body), 'B')
) STORED;

CREATE INDEX articles_fts_idx ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(
    fts,
    websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
) DESC;
```

### Dictionary Tuning

Penalty weights can be changed at runtime:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (w_deriv = 3.5, w_short_char = 100.0);
```

The default `script_mode = auto` detects supported modern Kazakh Latin orthography and returns Cyrillic stems. Disable Latin handling when strict Cyrillic-only behavior is required:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (script_mode = cyrillic_only);
```

### Upgrade and Search Caveats

- Stemmer upgrades change index terms. After upgrading to `2.3.0`, recompute stored `tsvector` columns or repopulate trigger-maintained vectors, then `VACUUM (ANALYZE)` the table.

```sql
ALTER EXTENSION pg_kazsearch UPDATE;
UPDATE articles SET title = title;
VACUUM (ANALYZE) articles;
```

- Long-lived sessions opened before an upgrade should reconnect so they load the new dictionary.
- Latin support targets the modern orthography. Mixed-script input, legacy apostrophe/acute/digraph spellings, and low-confidence ASCII tokens may remain unchanged.
- `websearch_to_tsquery` uses strict AND semantics for ordinary terms. Applications that need broader recall should deliberately implement and measure a fallback query rather than silently changing all searches to OR.
