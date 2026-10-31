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
| **2230** | {{< badge content="pg_fts" link="https://codeberg.org/gregburd/pg_fts" >}} | {{< ext "pg_fts" >}} | `1.15.1` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_rrf" >}} {{< ext "pgroonga" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] SQL 1.15.1 fixes configured-analyzer chunked reparsing at the 16,383-position HTML boundary; SQL update is a no-op and the format is unchanged. Known unfixed concurrent-flush ranked scans can duplicate rows. No routine REINDEX for this patch; indexes written by 1.10.0 or earlier still require rebuilding.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.15.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_fts` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.15.1` | {{< bg "18" "pg_fts_18" "green" >}} {{< bg "17" "pg_fts_17" "green" >}} {{< bg "16" "pg_fts_16" "red" >}} {{< bg "15" "pg_fts_15" "red" >}} {{< bg "14" "pg_fts_14" "red" >}} | `pg_fts_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.14.3` | {{< bg "18" "postgresql-18-pg-fts" "green" >}} {{< bg "17" "postgresql-17-pg-fts" "green" >}} {{< bg "16" "postgresql-16-pg-fts" "red" >}} {{< bg "15" "postgresql-15-pg-fts" "red" >}} {{< bg "14" "postgresql-14-pg-fts" "red" >}} | `postgresql-$v-pg-fts` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.15.1" "pg_fts_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_fts_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_fts_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-18-pg-fts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.14.3" "postgresql-17-pg-fts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-fts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-fts : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_18` | `1.15.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 602.9 KiB | [pg_fts_18-1.15.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_18-1.15.1-1PGSTY.el8.x86_64.rpm) |
| `pg_fts_18` | `1.15.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 589.5 KiB | [pg_fts_18-1.15.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_18-1.15.1-1PGSTY.el8.aarch64.rpm) |
| `pg_fts_18` | `1.15.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 575.6 KiB | [pg_fts_18-1.15.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_18-1.15.1-1PGSTY.el9.x86_64.rpm) |
| `pg_fts_18` | `1.15.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 569.4 KiB | [pg_fts_18-1.15.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_18-1.15.1-1PGSTY.el9.aarch64.rpm) |
| `pg_fts_18` | `1.15.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 583.5 KiB | [pg_fts_18-1.15.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_18-1.15.1-1PGSTY.el10.x86_64.rpm) |
| `pg_fts_18` | `1.15.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 574.3 KiB | [pg_fts_18-1.15.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_18-1.15.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-fts` | `1.14.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 569.3 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 555.3 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 572.8 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 558.1 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 589.2 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 583.5 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 568.3 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 562.5 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 564.6 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-fts` | `1.14.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 557.7 KiB | [postgresql-18-pg-fts_1.14.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-18-pg-fts_1.14.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fts_17` | `1.15.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 602.8 KiB | [pg_fts_17-1.15.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fts_17-1.15.1-1PGSTY.el8.x86_64.rpm) |
| `pg_fts_17` | `1.15.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 589.3 KiB | [pg_fts_17-1.15.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fts_17-1.15.1-1PGSTY.el8.aarch64.rpm) |
| `pg_fts_17` | `1.15.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 575.1 KiB | [pg_fts_17-1.15.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fts_17-1.15.1-1PGSTY.el9.x86_64.rpm) |
| `pg_fts_17` | `1.15.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 569.3 KiB | [pg_fts_17-1.15.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fts_17-1.15.1-1PGSTY.el9.aarch64.rpm) |
| `pg_fts_17` | `1.15.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 583.5 KiB | [pg_fts_17-1.15.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_fts_17-1.15.1-1PGSTY.el10.x86_64.rpm) |
| `pg_fts_17` | `1.15.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 574.2 KiB | [pg_fts_17-1.15.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_fts_17-1.15.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-fts` | `1.14.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 569.1 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 555.1 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 572.3 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 558.6 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 619.6 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 611.3 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 568.1 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 562.4 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 564.6 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-fts` | `1.14.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 557.5 KiB | [postgresql-17-pg-fts_1.14.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-fts/postgresql-17-pg-fts_1.14.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_fts" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_fts" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_fts-1.15.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_fts;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_fts;		# install via package name, for the active PG version

pig install pg_fts -v 18;   # install for PG 18
pig install pg_fts -v 17;   # install for PG 17

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_fts;
```

## Usage

Sources:

- [1.15.0 to 1.15.1 upgrade](https://api.pgxn.org/src/pg_fts/pg_fts-1.15.1/pg_fts--1.15.0--1.15.1.sql)
- [README.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.15.1/README.md)
- [pg_fts.control](https://api.pgxn.org/src/pg_fts/pg_fts-1.15.1/pg_fts.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.15.1/CHANGELOG.md)
- [pg_fts--1.15.1.sql](https://api.pgxn.org/src/pg_fts/pg_fts-1.15.1/pg_fts--1.15.1.sql)

`pg_fts` 1.15.1 provides full-text search with BM25/BM25F ranking and the dedicated fts inverted index. PostgreSQL 17 and 18 are supported; PostgreSQL 19/master CI is best-effort. The control is trusted and relocatable, and core use needs no shared preload.

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

### Upgrade to 1.10.0

After installing matching files, run `ALTER EXTENSION pg_fts UPDATE TO '1.10.0'`. Version 1.10.0 shares each segment's document-length array across backends by default and adds `fts_shared_doclen_stats()` to inspect the shared cache. A superuser can change `pg_fts.shared_doclen`; unavailable shared memory falls back to a private copy. There is no on-disk format change and no REINDEX requirement. Earlier delete-ranking and posting-block fixes remain included; distinguish shared-cache memory from fallback per-backend copies when planning capacity.

### Version 1.11.0

Version 1.11.0 fixes wrong-result cases and concurrent insert/VACUUM/merge failures. Ranked traversal uses best-first block ordering; `pg_fts.bestfirst` defaults on. The index format is unchanged and the SQL upgrade adds no objects. REINDEX is not required merely to upgrade, but it repairs index entries already left or lost by the earlier bugs.

### Version 1.12.0

Version 1.12.0 caches the segment dictionary index in each backend, charging it against `pg_fts.doclen_cache_mb`; a segment that does not fit falls back to the page-chain lookup. Changes to that budget now take effect without waiting for a segment-directory change. There is no index-format change, no new SQL object and no REINDEX requirement. A cancellation or error during cache reconstruction can leave memory allocated until that backend exits; this known limitation does not change query results.

### Compound Proximity and 1.13.0 Fixes

Queries now support unordered `w/N`, ordered `p/N`, exact-gap `<N>`, adjacency `<->`, and same-position `<0>`. Either operand can be a term, phrase, OR group or another proximity group; direct AND/NOT operands are rejected. Quoted phrases preserve removed-stopword gaps. Ranked prefix, fuzzy and regex searches now include all matching rows, and pending documents participate in ranking; expansion-only matches can still score zero because scores use literal query terms.

This release fixes malformed binary-query crashes, non-ASCII fuzzy matching and a positional-scan double free during concurrent merges. Install matching files and run `ALTER EXTENSION pg_fts UPDATE TO '1.13.0'`; no index-format change or routine REINDEX is required.

### 1.14.0 Vacuum and Write-Load Fixes

Version 1.14.0 bounds posting-list vacuum batches by `maintenance_work_mem` and skips tail truncation during parallel VACUUM. Deferred free space remains reusable and can be reclaimed by a later non-parallel VACUUM or `fts_vacuum()`. The new `fts_matchsel` estimator improves ranked-search plan selection. `pg_fts.pending_flush_threshold` defaults to 20000 pending documents; 0 disables insert-triggered flushing. The insert path attempts the maintenance lock without waiting.

Install matching files and run `ALTER EXTENSION pg_fts UPDATE TO '1.14.0'` to rewire the match operators. Index format is unchanged; routine REINDEX is unnecessary.

### Counting and Maintenance Fixes in 1.14.1

1.14.1 prevents the exact-count fast path from using partial, invalid or snapshot-ineligible new indexes, and corrects selectivity estimates inflated by pending documents. VACUUM no longer evacuates tail pages when it cannot acquire the truncation lock, avoiding index growth under concurrent readers; bounded memory batches also cover configurations without the document-length sidecar. Install matching files and run `ALTER EXTENSION pg_fts UPDATE TO '1.14.1'`. The upgrade from 1.14.0 changes neither SQL objects nor index format and needs no routine REINDEX.

### 1.14.2: Rebuild Indexes Written Before 1.11.0

Version 1.14.2 changes documentation and the extension version only; its SQL upgrade is a no-op. Rebuild every index that received writes under 1.10.0 or earlier with REINDEX. Those versions could leave false matches after aborted or dead pending inserts, and omit rows inserted during a flush. An in-place extension update, VACUUM, fts_merge or fts_vacuum cannot repair that old state. A clean result from probing a few terms is not proof that the index is sound. Indexes built and written only by 1.11.0 or later do not need this historical repair.

### 1.14.3 Fixes and Remaining Race

Version 1.14.3 fixes a ranked-scan backend crash during frequent concurrent pending-list flushes and an out-of-bounds read when a BM25 scorer receives a short document-frequency array. It also corrects pending-document frequencies, partial-index plan estimates and some boolean-query work. Install matching files, replace old-library sessions and run the no-op SQL upgrade to 1.14.3. Formats and SQL objects are unchanged; the historical REINDEX requirement above still applies.

A separate concurrent-flush race remains unfixed: ranked scans can return a pending row twice and double-count ranking statistics. Bitmap scans, count queries and match filters remove duplicate tuple IDs and are unaffected by that specific race. Do not assume this patch makes ranked results duplicate-free under concurrent flushes.

### 1.15.1 Configured-Analyzer Repair

Version 1.15.1 fixes configured-analyzer chunked reparsing that could split HTML markup at the 16,383-position boundary. It changes neither the on-disk index format nor SQL objects, and its SQL upgrade is a no-op. For documents written by 1.15.0 whose markup crosses that boundary, recompute the stored `ftsdoc` and REINDEX the expression index. Data from 1.14.3 or earlier still needs the historical repair described in the 1.15.0 upgrade notice.
