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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `zhparser` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "zhparser_18" "green" >}} {{< bg "17" "zhparser_17" "green" >}} {{< bg "16" "zhparser_16" "green" >}} {{< bg "15" "zhparser_15" "green" >}} {{< bg "14" "zhparser_14" "green" >}} | `zhparser_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "postgresql-18-zhparser" "green" >}} {{< bg "17" "postgresql-17-zhparser" "green" >}} {{< bg "16" "postgresql-16-zhparser" "green" >}} {{< bg "15" "postgresql-15-zhparser" "green" >}} {{< bg "14" "postgresql-14-zhparser" "green" >}} | `postgresql-$v-zhparser` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.3" "zhparser_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "zhparser_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-zhparser : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-zhparser : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-zhparser : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-zhparser : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-zhparser : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION zhparser;
```



## Usage

> [GitHub: amutu/zhparser](https://github.com/amutu/zhparser)

`zhparser` is a PostgreSQL extension for full-text search of Chinese, based on the Simple Chinese Word Segmentation (SCWS) library.

## Features

- Chinese text segmentation for PostgreSQL full-text search
- Built on the SCWS (Simple Chinese Word Segmentation) library
- Supports custom dictionaries (TXT and XDB formats)
- Database-level custom word tables (since v2.1)
- Multiple tunable parameters for segmentation behavior

## Quick Start

```sql
-- Create the extension
CREATE EXTENSION zhparser;

-- Create a text search configuration using zhparser
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);

-- Add token type mappings
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;

-- Test Chinese text segmentation
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所，后在日本京都大学深造');

-- Create a table and index for Chinese full text search
CREATE TABLE articles (id serial PRIMARY KEY, title text, body text);

CREATE INDEX articles_body_idx ON articles
  USING gin (to_tsvector('chinese', body));

-- Query with Chinese full text search
SELECT * FROM articles
  WHERE to_tsvector('chinese', body) @@ to_tsquery('chinese', '中国');
```

## Configuration Parameters

zhparser provides several GUC parameters to control segmentation behavior:

| Parameter | Default | Description |
|-----------|---------|-------------|
| `zhparser.punctuation_ignore` | `off` | Ignore all punctuation |
| `zhparser.seg_with_duality` | `off` | Perform duality segmentation on long words |
| `zhparser.dict_in_memory` | `off` | Load the whole dictionary into memory |
| `zhparser.multi_short` | `off` | Short word compound segmentation |
| `zhparser.multi_duality` | `off` | Duality compound segmentation |
| `zhparser.multi_zmain` | `off` | Key word in first compound segmentation |
| `zhparser.multi_zall` | `off` | Use all compound segmentation |

## Token Types

zhparser supports the following token types from SCWS:

| Code | Description |
|------|-------------|
| `a` | Adjective |
| `b` | Differentiation (区别词) |
| `c` | Conjunction |
| `d` | Adverb |
| `e` | Exclamation |
| `f` | Position word (方位词) |
| `g` | Root word (词根) |
| `h` | Prefix |
| `i` | Idiom |
| `j` | Abbreviation |
| `k` | Suffix |
| `l` | Temporary idiom |
| `m` | Numeral |
| `n` | Noun |
| `o` | Onomatopoeia |
| `p` | Preposition |
| `q` | Classifier |
| `r` | Pronoun |
| `s` | Space word (处所词) |
| `t` | Time word |
| `u` | Auxiliary |
| `v` | Verb |
| `w` | Punctuation |
| `x` | Unknown |
| `y` | Modal particle |
| `z` | Status word (状态词) |

## Custom Dictionaries

### File-based Dictionaries

Place custom dictionary files in the share directory (typically `$SHAREDIR/tsearch_data/`):

- TXT format: one word per line
- XDB format: compiled SCWS dictionary format

Custom dictionaries take precedence over built-in dictionaries.

### Database-level Custom Words (v2.1+)

```sql
-- Add custom words via zhparser's built-in table
INSERT INTO zhparser.zhprs_custom_word VALUES ('中国科学院计算所');

-- Reload custom dictionary (reconnect after sync to take effect)
SELECT sync_zhprs_custom_word();

-- Verify segmentation with custom word
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所');
```

## Docker Quick Start

```bash
docker run --name pgzhparser -d \
  -e POSTGRES_PASSWORD=somepassword \
  zhparser/zhparser:bookworm-16
```
