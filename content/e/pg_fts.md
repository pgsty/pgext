---
title: "pg_fts"
linkTitle: "pg_fts"
description: "Full-text search with BM25 and BM25F ranking"
weight: 2230
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_fts**](https://codeberg.org/gregburd/pg_fts) : Full-text search with BM25 and BM25F ranking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2230** | {{< badge content="pg_fts" link="https://codeberg.org/gregburd/pg_fts" >}} | {{< ext "pg_fts" >}} | `1.9.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_rrf" >}} {{< ext "pgroonga" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] PG17-18; trusted and relocatable.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.9.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_fts` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.9.0` | {{< bg "18" "pg_fts_18" "green" >}} {{< bg "17" "pg_fts_17" "green" >}} {{< bg "16" "pg_fts_16" "red" >}} {{< bg "15" "pg_fts_15" "red" >}} {{< bg "14" "pg_fts_14" "red" >}} | `pg_fts_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.9.0` | {{< bg "18" "postgresql-18-pg-fts" "green" >}} {{< bg "17" "postgresql-17-pg-fts" "green" >}} {{< bg "16" "postgresql-16-pg-fts" "red" >}} {{< bg "15" "postgresql-15-pg-fts" "red" >}} {{< bg "14" "postgresql-14-pg-fts" "red" >}} | `postgresql-$v-pg-fts` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.9.0" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_18` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 469.1 KiB | [pg_fts_18-1.9.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_18-1.9.0-1PGSTY.el8.x86_64.rpm) |
| `pg_fts_18` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 458.9 KiB | [pg_fts_18-1.9.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_18-1.9.0-1PGSTY.el8.aarch64.rpm) |
| `pg_fts_18` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 446.0 KiB | [pg_fts_18-1.9.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_18-1.9.0-1PGSTY.el9.x86_64.rpm) |
| `pg_fts_18` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 440.3 KiB | [pg_fts_18-1.9.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_18-1.9.0-1PGSTY.el9.aarch64.rpm) |
| `pg_fts_18` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 450.9 KiB | [pg_fts_18-1.9.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_18-1.9.0-1PGSTY.el10.x86_64.rpm) |
| `pg_fts_18` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 442.9 KiB | [pg_fts_18-1.9.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_18-1.9.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-fts` | `1.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 453.5 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 442.0 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 455.6 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 444.4 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 469.0 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 462.0 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 448.4 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 443.8 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 446.0 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-fts` | `1.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 439.1 KiB | [postgresql-18-pg-fts_1.9.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_1.9.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_17` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 469.2 KiB | [pg_fts_17-1.9.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_17-1.9.0-1PGSTY.el8.x86_64.rpm) |
| `pg_fts_17` | `1.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 458.9 KiB | [pg_fts_17-1.9.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_17-1.9.0-1PGSTY.el8.aarch64.rpm) |
| `pg_fts_17` | `1.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 446.0 KiB | [pg_fts_17-1.9.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_17-1.9.0-1PGSTY.el9.x86_64.rpm) |
| `pg_fts_17` | `1.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 440.3 KiB | [pg_fts_17-1.9.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_17-1.9.0-1PGSTY.el9.aarch64.rpm) |
| `pg_fts_17` | `1.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 450.9 KiB | [pg_fts_17-1.9.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_17-1.9.0-1PGSTY.el10.x86_64.rpm) |
| `pg_fts_17` | `1.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 442.9 KiB | [pg_fts_17-1.9.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_17-1.9.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-fts` | `1.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 453.5 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 442.0 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 455.6 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 444.4 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 495.8 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 486.9 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 448.3 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 444.0 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 445.9 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-fts` | `1.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 439.2 KiB | [postgresql-17-pg-fts_1.9.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_1.9.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_fts" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_fts" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_fts-1.9.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_fts;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

- [README.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/README.md)
- [pg_fts.control](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts.control)
- [CHANGELOG.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/CHANGELOG.md)
- [pg_fts--1.8.6--1.9.0.sql](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts--1.8.6--1.9.0.sql)

`pg_fts` 1.9.0 provides full-text search with BM25/BM25F ranking and the dedicated fts inverted index. PostgreSQL 17 and 18 are supported; PostgreSQL 19/master CI is best-effort. The control is trusted and relocatable, and core use needs no shared preload.

### Search and Rank

```sql
CREATE EXTENSION pg_fts;
CREATE TABLE docs (id bigint, body text);
CREATE INDEX docs_fts ON docs USING fts (to_ftsdoc('english', body));
SELECT id FROM docs
 WHERE to_ftsdoc('english', body) @@@ to_ftsquery('english', 'quick fox')
 ORDER BY to_ftsdoc('english', body) <=> to_ftsquery('english', 'quick fox')
 LIMIT 10;
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

### Objects and Queries

`ftsdoc` and `ftsquery` represent analyzed documents and queries. `to_ftsdoc()` and `to_ftsquery()` construct them; `@@@` matches documents and `<=>` orders relevance distance. The match predicate is required for an indexed KNN ordering scan. Queries support boolean terms, phrases, prefix, fuzzy and regular-expression matching. Multi-column documents support BM25F field weighting. `fts_count()` and ordinary count queries can use exact index counting; `fts_search()` exposes direct ranked results. Regex/long-fuzzy acceleration is opt-in through `trigrams = on`.

### Maintenance and Privileges

Pending inserts are immediately searchable; merging is not a visibility prerequisite. `fts_merge()` compacts segments and `fts_vacuum()` reclaims physical space. Both write WAL, require index ownership and must run on a primary. Large pending documents can temporarily consume substantial space, so merge and reclaim periodically during bulk ingestion. `fts_search()` and `fts_anomalous_docs()` expose indexed content and are revoked from PUBLIC by default; widen access only deliberately. Regular table-query visibility and direct helper access are different permission surfaces.

### Upgrade to 1.9.0

After installing matching files, run `ALTER EXTENSION pg_fts UPDATE TO '1.9.0'`. This release fixes incorrect ranking after deletes and missed top-k results at posting-block boundaries. There is no on-disk format change and no REINDEX requirement. `pg_fts.doclen_cache_mb` defaults to 64 (0 disables the cache); `pg_fts.dense_score_min_df` defaults to 32768 (0 disables that scoring path). Account for per-backend cache memory.
