---
title: "zhparser"
linkTitle: "zhparser"
description: "a parser for full-text search of Chinese"
weight: 2130
categories: ["FTS"]
width: full
---

[**zhparser**](https://github.com/amutu/zhparser) : a parser for full-text search of Chinese


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2130** | {{< badge content="zhparser" link="https://github.com/amutu/zhparser" >}} | {{< ext "zhparser" >}} | `2.3` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pgroonga_database" >}} {{< ext "pg_bigm" >}} {{< ext "pg_tokenizer" >}} {{< ext "vchord_bm25" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `zhparser` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "zhparser_18" "green" >}} {{< bg "17" "zhparser_17" "green" >}} {{< bg "16" "zhparser_16" "green" >}} {{< bg "15" "zhparser_15" "green" >}} {{< bg "14" "zhparser_14" "green" >}} {{< bg "13" "zhparser_13" "green" >}} | `zhparser_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "postgresql-18-zhparser" "green" >}} {{< bg "17" "postgresql-17-zhparser" "green" >}} {{< bg "16" "postgresql-16-zhparser" "green" >}} {{< bg "15" "postgresql-15-zhparser" "green" >}} {{< bg "14" "postgresql-14-zhparser" "green" >}} {{< bg "13" "postgresql-13-zhparser" "green" >}} | `postgresql-$v-zhparser` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-13-zhparser : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_18` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_18-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_18-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_18` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_18-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_18-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_18` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_18-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_18` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_18-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_18` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_18-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_18` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_18-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_18-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-18-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-18-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_17` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_17-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_17-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_17` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_17-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_17-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_17` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_17-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_17` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_17-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_17` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_17-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_17` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_17-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_17-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-17-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-17-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_16` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_16-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_16-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_16` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_16-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_16-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_16` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_16-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_16` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_16-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_16` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_16-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_16` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_16-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_16-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-16-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-16-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_15` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_15-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_15-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_15` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_15-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_15-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_15` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_15-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_15` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_15-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_15` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_15-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_15` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_15-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_15-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-15-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-15-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_14` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_14-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_14-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_14` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_14-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_14-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_14` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_14-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_14` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_14-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_14` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_14-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_14` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_14-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_14-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-14-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-14-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `zhparser_13` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.7 MiB | [zhparser_13-2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/zhparser_13-2.3-1PIGSTY.el8.x86_64.rpm) |
| `zhparser_13` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.7 MiB | [zhparser_13-2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/zhparser_13-2.3-1PIGSTY.el8.aarch64.rpm) |
| `zhparser_13` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/zhparser_13-2.3-1PIGSTY.el9.x86_64.rpm) |
| `zhparser_13` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/zhparser_13-2.3-1PIGSTY.el9.aarch64.rpm) |
| `zhparser_13` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/zhparser_13-2.3-1PIGSTY.el10.x86_64.rpm) |
| `zhparser_13` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.3 MiB | [zhparser_13-2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/zhparser_13-2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-zhparser` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.0 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-zhparser` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.0 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-zhparser` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.0 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-zhparser` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.0 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-zhparser` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-zhparser` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-zhparser` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-zhparser` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-13-zhparser_2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/z/zhparser/postgresql-13-zhparser_2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/amutu/zhparser" title="Repository" icon="github" subtitle="github.com/amutu/zhparser" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="zhparser-2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg zhparser;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install zhparser;		# install via package name, for the active PG version

pig install zhparser -v 18;   # install for PG 18
pig install zhparser -v 17;   # install for PG 17
pig install zhparser -v 16;   # install for PG 16
pig install zhparser -v 15;   # install for PG 15
pig install zhparser -v 14;   # install for PG 14
pig install zhparser -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION zhparser;
```
