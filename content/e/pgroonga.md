---
title: "pgroonga"
linkTitle: "pgroonga"
description: "Use Groonga as index, fast full text search platform for all languages!"
weight: 2110
categories: ["FTS"]
width: full
---

[**pgroonga**](https://github.com/pgroonga/pgroonga) : Use Groonga as index, fast full text search platform for all languages!


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2110** | {{< badge content="pgroonga" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga" >}} | `4.0.4` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_tokenizer" >}} {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "vchord_bm25" >}} |
|    **Siblings**   | {{< ext "pgroonga_database" >}} |

> [!Note] require xxHash vendor repo to build


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgroonga` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "pgroonga_18*" "green" >}} {{< bg "17" "pgroonga_17*" "green" >}} {{< bg "16" "pgroonga_16*" "green" >}} {{< bg "15" "pgroonga_15*" "green" >}} {{< bg "14" "pgroonga_14*" "green" >}} {{< bg "13" "pgroonga_13*" "green" >}} | `pgroonga_$v*` | `groonga-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.4` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} {{< bg "13" "postgresql-13-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "pgroonga_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-14-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.4" "postgresql-13-pgroonga : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_18` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 360.6 KiB | [pgroonga_18-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_18-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_18` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 348.8 KiB | [pgroonga_18-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_18-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_18` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 345.3 KiB | [pgroonga_18-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_18-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_18` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 337.0 KiB | [pgroonga_18-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_18-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_18` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 347.4 KiB | [pgroonga_18-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_18-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_18` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 339.4 KiB | [pgroonga_18-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_18-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-18-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 621.9 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 612.8 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 621.9 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 613.4 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 678.1 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 680.2 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 651.0 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 650.1 KiB | [postgresql-18-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-18-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_17` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 360.3 KiB | [pgroonga_17-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_17-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_17` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 348.4 KiB | [pgroonga_17-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_17-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_17` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 345.3 KiB | [pgroonga_17-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_17-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_17` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 337.3 KiB | [pgroonga_17-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_17-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_17` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 347.1 KiB | [pgroonga_17-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_17-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_17` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 339.3 KiB | [pgroonga_17-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_17-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-17-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 621.4 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 612.1 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 621.7 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 612.9 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 757.2 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 759.0 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 650.0 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 649.4 KiB | [postgresql-17-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-17-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_16` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 357.8 KiB | [pgroonga_16-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_16-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_16` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 346.3 KiB | [pgroonga_16-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_16-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_16` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 342.6 KiB | [pgroonga_16-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_16-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_16` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 335.5 KiB | [pgroonga_16-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_16-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_16` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.9 KiB | [pgroonga_16-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_16-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_16` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 337.1 KiB | [pgroonga_16-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_16-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-16-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 615.0 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 606.8 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 615.2 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 607.0 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 744.3 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 746.9 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 643.0 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 643.0 KiB | [postgresql-16-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-16-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_15` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 360.6 KiB | [pgroonga_15-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_15-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_15` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.2 KiB | [pgroonga_15-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_15-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_15` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 346.6 KiB | [pgroonga_15-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_15-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_15` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 339.9 KiB | [pgroonga_15-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_15-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_15` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 349.2 KiB | [pgroonga_15-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_15-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_15` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 339.4 KiB | [pgroonga_15-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_15-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-15-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 617.3 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 608.1 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 618.9 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 608.2 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 751.6 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 758.6 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 650.9 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 651.3 KiB | [postgresql-15-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-15-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_14` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.1 KiB | [pgroonga_14-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_14-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_14` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 332.6 KiB | [pgroonga_14-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_14-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_14` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 328.2 KiB | [pgroonga_14-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_14-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_14` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 322.9 KiB | [pgroonga_14-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_14-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_14` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 331.0 KiB | [pgroonga_14-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_14-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_14` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 322.9 KiB | [pgroonga_14-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_14-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-14-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 565.5 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 558.2 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 566.9 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 559.0 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 690.8 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 698.0 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 596.8 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 598.6 KiB | [postgresql-14-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-14-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgroonga_13` | `4.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 340.2 KiB | [pgroonga_13-4.0.4-1.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgroonga_13-4.0.4-1.el8.x86_64.rpm) |
| `pgroonga_13` | `4.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 332.6 KiB | [pgroonga_13-4.0.4-1.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgroonga_13-4.0.4-1.el8.aarch64.rpm) |
| `pgroonga_13` | `4.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 327.8 KiB | [pgroonga_13-4.0.4-1.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgroonga_13-4.0.4-1.el9.x86_64.rpm) |
| `pgroonga_13` | `4.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 323.1 KiB | [pgroonga_13-4.0.4-1.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgroonga_13-4.0.4-1.el9.aarch64.rpm) |
| `pgroonga_13` | `4.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.4 KiB | [pgroonga_13-4.0.4-1.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgroonga_13-4.0.4-1.el10.x86_64.rpm) |
| `pgroonga_13` | `4.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 322.9 KiB | [pgroonga_13-4.0.4-1.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgroonga_13-4.0.4-1.el10.aarch64.rpm) |
| `postgresql-13-pgroonga` | `4.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 667.1 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 655.8 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 658.1 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 645.2 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 727.9 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 730.3 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 698.2 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgroonga` | `4.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 697.5 KiB | [postgresql-13-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgroonga/postgresql-13-pgroonga_4.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgroonga;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgroonga;		# install via package name, for the active PG version

pig install pgroonga -v 18;   # install for PG 18
pig install pgroonga -v 17;   # install for PG 17
pig install pgroonga -v 16;   # install for PG 16
pig install pgroonga -v 15;   # install for PG 15
pig install pgroonga -v 14;   # install for PG 14
pig install pgroonga -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgroonga;
```


## Usage

- https://pgroonga.github.io/
- [News](https://pgroonga.github.io/news/): It lists release information.
- [Overview](https://pgroonga.github.io/overview/): It describes about PGroonga.
- [Install](https://pgroonga.github.io/install/): It describes how to install PGroonga.
- [Upgrade](https://pgroonga.github.io/upgrade/): It describes how to upgrade PGroonga.
- [Uninstall](https://pgroonga.github.io/uninstall/): It describes how to uninstall PGroonga.
- [Tutorial](https://pgroonga.github.io/tutorial/): It describes how to use PGroonga step by step.
- [FAQ](https://pgroonga.github.io/faq/): Frequently asked questions.
- [How to](https://pgroonga.github.io/how-to/): It describes about useful information for specific situations.
- [Reference](https://pgroonga.github.io/reference/): It describes details for each features such as options, functions and operators.
- [Troubleshooting](https://pgroonga.github.io/troubleshooting/): It describes how to fix troubles.
- [Community](https://pgroonga.github.io/community/): It introduces about PGroonga community.
- [Users](https://pgroonga.github.io/users/): It lists PGroonga users.
- [Development](https://pgroonga.github.io/development/): It describes how to develop PGroonga.

Here's a quick [tutorial](https://pgroonga.github.io/tutorial/) about how to use PGroonga:

```sql
CREATE EXTENSION IF NOT EXISTS pgroonga;

CREATE TABLE memos
(
    id      integer,
    content text
);

CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as index.');
INSERT INTO memos VALUES (4, 'There is groonga command.');

SET enable_seqscan = off;

-- now let's query pgroonga

SELECT * FROM memos WHERE content &@ 'engine';
--  id |                                content                                 
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)

SELECT * FROM memos WHERE content &@~ 'PGroonga OR PostgreSQL';
--  id |                            content                             
-- ----+----------------------------------------------------------------
--   3 | PGroonga is a PostgreSQL extension that uses Groonga as index.
--   1 | PostgreSQL is a relational database management system.
-- (2 rows)

SELECT * FROM memos WHERE content LIKE '%engine%';
--  id |                                content                                 
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)
```

