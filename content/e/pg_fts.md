---
title: "pg_fts"
linkTitle: "pg_fts"
description: "Full-text search with BM25 and BM25F ranking"
weight: 2220
categories: ["FTS"]
width: full
---

[**pg_fts**](https://codeberg.org/gregburd/pg_fts) : Full-text search with BM25 and BM25F ranking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2220** | {{< badge content="pg_fts" link="https://codeberg.org/gregburd/pg_fts" >}} | {{< ext "pg_fts" >}} | `0.2.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "vchord_bm25" >}} {{< ext "psql_bm25s" >}} {{< ext "pg_bestmatch" >}} |

> [!Note] Requires PostgreSQL 17 or newer; the control file marks the extension trusted and relocatable; RPM builds also provide an llvmjit subpackage.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_fts` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pg_fts_18" "green" >}} {{< bg "17" "pg_fts_17" "green" >}} {{< bg "16" "pg_fts_16" "red" >}} {{< bg "15" "pg_fts_15" "red" >}} {{< bg "14" "pg_fts_14" "red" >}} | `pg_fts_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pg-fts" "green" >}} {{< bg "17" "postgresql-17-pg-fts" "green" >}} {{< bg "16" "postgresql-16-pg-fts" "red" >}} {{< bg "15" "postgresql-15-pg-fts" "red" >}} {{< bg "14" "postgresql-14-pg-fts" "red" >}} | `postgresql-$v-pg-fts` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_fts_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_fts_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_fts_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-fts : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-fts : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 113.6 KiB | [pg_fts_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fts_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 109.9 KiB | [pg_fts_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fts_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.3 KiB | [pg_fts_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fts_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 88.3 KiB | [pg_fts_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fts_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.4 KiB | [pg_fts_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fts_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.3 KiB | [pg_fts_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-fts` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 265.4 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 258.2 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 266.9 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 260.3 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 271.9 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 268.8 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 256.4 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 253.2 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 256.5 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-fts` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 252.6 KiB | [postgresql-18-pg-fts_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_0.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 113.6 KiB | [pg_fts_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fts_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 109.9 KiB | [pg_fts_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fts_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.3 KiB | [pg_fts_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fts_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 88.3 KiB | [pg_fts_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fts_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.4 KiB | [pg_fts_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_fts_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.2 KiB | [pg_fts_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-fts` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 265.5 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 258.1 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 266.8 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 260.2 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 290.0 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 285.8 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 256.4 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 253.1 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 256.5 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-fts` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 252.5 KiB | [postgresql-17-pg-fts_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_0.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_fts" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_fts" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_fts-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_fts;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_fts;		# install via package name, for the active PG version

pig install pg_fts -v 18;   # install for PG 18
pig install pg_fts -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_fts;
```

## Usage

Sources:

- [Official v0.2.0 README](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/README.md)
- [v0.2.0 changelog](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/CHANGELOG.md)
- [v0.2.0 SQL API](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts--0.2.0.sql)
- [v0.2.0 control file](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts.control)

`pg_fts` provides BM25/BM25F full-text ranking through dedicated `ftsdoc` and `ftsquery` types and an `fts` inverted-index access method. It supports boolean, phrase, NEAR, prefix, fuzzy, and regular-expression terms while keeping corpus statistics in the index for relevance scoring. Version `0.2.0` requires PostgreSQL 17 or newer.

### Create and Query an Index

```sql
CREATE EXTENSION pg_fts;

CREATE TABLE docs (
    id bigint PRIMARY KEY,
    body text NOT NULL
);

CREATE INDEX docs_fts
ON docs USING fts (to_ftsdoc('english', body));
```

Use the same text-search configuration for documents and ordinary query terms:

```sql
WITH q AS (
    SELECT to_ftsquery('english', 'postgres & "query planner" & index*') AS query
)
SELECT d.id,
       fts_snippet(d.body, q.query) AS excerpt
FROM docs AS d
CROSS JOIN q
WHERE to_ftsdoc('english', d.body) @@@ q.query
ORDER BY to_ftsdoc('english', d.body) <=> q.query
LIMIT 10;
```

`@@@` matches, while ascending `<=>` distance orders rows by descending relevance and can drive an index ordering scan for top-k queries.

### Query Language and API Index

- `to_ftsdoc([regconfig,] text)` and `to_ftsquery([regconfig,] text)`: analyze documents and parse queries.
- `quick brown`, `quick & brown`, `quick | brown`, and `!slow`: implicit/explicit AND, OR, and NOT.
- `"quick brown"`, `NEAR(...)`, `term*`, `term~2`, and `/regular-expression/`: phrase, proximity, prefix, fuzzy, and regex terms.
- `fts_bm25`, `fts_bm25_opts`, and `fts_bm25f`: explicit BM25 scoring variants and multi-field scoring.
- `fts_index_stats(index)` and `fts_index_df(index, query)`: index-maintained document count, average length, vocabulary size, and term frequencies.
- `fts_highlight` and `fts_snippet`: present matching text.
- `fts_search(index, query, k)` and `fts_count(index, query)`: index-native top-k and MVCC-aware count operations.
- `tsquery_to_ftsquery(tsquery)`: migration helper; it does not make `pg_fts` a transparent replacement for `tsvector`/GIN.

### Maintenance and Version Caveats

```sql
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

- Inserts enter an immediately matchable pending list, but ranked `<=>` and `fts_search` results cover merged segments. Run `fts_merge()` when newly inserted documents must participate in ranking immediately.
- `fts_vacuum()` compacts segments and truncates reclaimable index pages; ordinary `VACUUM` also participates in pending-list and tombstone maintenance.
- Version `0.2.0` renamed the access method from `bm25` to `fts`. Indexes created by `0.1.0` with `USING bm25` must be recreated.
- If the library reports an on-disk format mismatch, follow its `REINDEX` hint rather than attempting to read the index with a different format version.
- The access method is non-covering and does not provide parallel scans in this release. Provision the extension and index separately on logical-replication subscribers; indexes themselves are not logically replicated.

