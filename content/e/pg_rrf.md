---
title: "pg_rrf"
linkTitle: "pg_rrf"
description: "Reciprocal rank fusion functions for hybrid search"
weight: 1845
categories: ["RAG"]
width: full
---

[**pg_rrf**](https://github.com/yuiseki/pg_rrf) : Reciprocal rank fusion functions for hybrid search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1845** | {{< badge content="pg_rrf" link="https://github.com/yuiseki/pg_rrf" >}} | {{< ext "pg_rrf" >}} | `0.0.3` | {{< category "RAG" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_rrf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "pg_rrf_18" "red" >}} {{< bg "17" "pg_rrf_17" "green" >}} {{< bg "16" "pg_rrf_16" "green" >}} {{< bg "15" "pg_rrf_15" "green" >}} {{< bg "14" "pg_rrf_14" "green" >}} | `pg_rrf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "postgresql-18-pg-rrf" "red" >}} {{< bg "17" "postgresql-17-pg-rrf" "green" >}} {{< bg "16" "postgresql-16-pg-rrf" "green" >}} {{< bg "15" "postgresql-15-pg-rrf" "green" >}} {{< bg "14" "postgresql-14-pg-rrf" "green" >}} | `postgresql-$v-pg-rrf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_rrf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-rrf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_17` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.3 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 209.0 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.9 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 222.9 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.3 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_17-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 222.8 KiB | [pg_rrf_17-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_17-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 256.8 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 159.8 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 256.5 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 159.8 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 290.2 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 184.9 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 287.8 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.7 KiB | [postgresql-17-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_16` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.4 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 208.9 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.8 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.0 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.4 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_16-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 222.9 KiB | [pg_rrf_16-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_16-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 257.1 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 159.8 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 256.9 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 159.8 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 290.2 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 184.7 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 287.9 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.6 KiB | [postgresql-16-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_15` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.0 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 209.0 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.7 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 222.9 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.2 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_15-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 222.9 KiB | [pg_rrf_15-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_15-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 256.8 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 159.8 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 256.8 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 159.9 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 290.0 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 184.7 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 287.7 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.7 KiB | [postgresql-15-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_14` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.1 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 209.0 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.8 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.1 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 330.0 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_14-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.0 KiB | [pg_rrf_14-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_14-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 256.6 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 159.8 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 256.8 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 159.9 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 289.8 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 184.7 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 287.4 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 183.6 KiB | [postgresql-14-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/yuiseki/pg_rrf" title="Repository" icon="github" subtitle="github.com/yuiseki/pg_rrf" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_rrf-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_rrf;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_rrf;		# install via package name, for the active PG version

pig install pg_rrf -v 17;   # install for PG 17
pig install pg_rrf -v 16;   # install for PG 16
pig install pg_rrf -v 15;   # install for PG 15
pig install pg_rrf -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_rrf;
```

## Usage
> Sources: [README](https://github.com/yuiseki/pg_rrf/blob/main/README.md) and [project repo](https://github.com/yuiseki/pg_rrf).

`pg_rrf` provides Reciprocal Rank Fusion functions for hybrid search score fusion.
It is focused on combining ranked candidate lists without hand-written `FULL OUTER JOIN` / `COALESCE` plumbing.

### Core Functions

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

The README also documents the behavior of the score helpers:

- missing ranks are ignored
- ranks `<= 0` are ignored
- `k <= 0` raises an error

### Example

```sql
CREATE EXTENSION pg_rrf;

SELECT rrf(1, 2, 60) AS rrf_12;
SELECT rrf3(1, 2, 3, 60) AS rrf_123;
SELECT rrfn(ARRAY[1, 2, 3], 60) AS rrfn_123;
SELECT *
FROM rrf_fuse(ARRAY[10, 20, 30], ARRAY[20, 40], 60)
ORDER BY score DESC;
```

### Hybrid Search Pattern

The upstream README shows `rrf_fuse` as a replacement for a manual fusion query:

```sql
WITH fused AS (
  SELECT *
  FROM rrf_fuse(
    ARRAY(SELECT id FROM docs ORDER BY bm25_score DESC LIMIT 100),
    ARRAY(SELECT id FROM docs ORDER BY embedding <=> :qvec LIMIT 100),
    60
  )
)
SELECT d.*, fused.score
FROM fused
JOIN docs d USING (id)
ORDER BY fused.score DESC
LIMIT 20;
```

### Requirements

- PostgreSQL 14-17
- Docker and Docker Compose v2

The README says the build and test flow runs in Docker, so local PostgreSQL and Rust toolchains are not required for the package workflow.
