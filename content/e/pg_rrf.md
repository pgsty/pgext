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


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_rrf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "pg_rrf_18" "green" >}} {{< bg "17" "pg_rrf_17" "green" >}} {{< bg "16" "pg_rrf_16" "green" >}} {{< bg "15" "pg_rrf_15" "green" >}} {{< bg "14" "pg_rrf_14" "green" >}} | `pg_rrf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "postgresql-18-pg-rrf" "green" >}} {{< bg "17" "postgresql-17-pg-rrf" "green" >}} {{< bg "16" "postgresql-16-pg-rrf" "green" >}} {{< bg "15" "postgresql-15-pg-rrf" "green" >}} {{< bg "14" "postgresql-14-pg-rrf" "green" >}} | `postgresql-$v-pg-rrf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_rrf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-rrf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-rrf : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_18` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 856.7 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_18-0.0.3-3PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_18` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 769.0 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_18-0.0.3-3PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_18` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 867.4 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_18-0.0.3-3PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_18` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 815.4 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_18-0.0.3-3PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_18` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 867.2 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_18-0.0.3-3PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_18` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 794.1 KiB | [pg_rrf_18-0.0.3-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_18-0.0.3-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 682.5 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 569.4 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 682.9 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 569.2 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 760.3 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 670.8 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 752.3 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 662.5 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 747.8 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-rrf` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 661.5 KiB | [postgresql-18-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-18-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_17` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 854.1 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_17-0.0.3-3PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 765.9 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_17-0.0.3-3PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 863.8 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_17-0.0.3-3PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 811.8 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_17-0.0.3-3PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 863.6 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_17-0.0.3-3PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_17` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 793.3 KiB | [pg_rrf_17-0.0.3-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_17-0.0.3-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 681.5 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 567.3 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 680.8 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 568.1 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 757.5 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 668.3 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 749.3 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 660.6 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 745.7 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-rrf` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 659.2 KiB | [postgresql-17-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-17-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_16` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 853.2 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_16-0.0.3-3PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 764.4 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_16-0.0.3-3PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 862.1 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_16-0.0.3-3PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 810.3 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_16-0.0.3-3PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 863.3 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_16-0.0.3-3PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_16` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 792.7 KiB | [pg_rrf_16-0.0.3-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_16-0.0.3-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 680.9 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.7 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 681.3 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 567.0 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 758.0 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 667.8 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 749.9 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 659.6 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 745.3 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-rrf` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 658.6 KiB | [postgresql-16-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-16-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_15` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 842.9 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_15-0.0.3-3PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 755.1 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_15-0.0.3-3PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 852.2 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_15-0.0.3-3PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 800.6 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_15-0.0.3-3PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 852.5 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_15-0.0.3-3PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_15` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 787.9 KiB | [pg_rrf_15-0.0.3-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_15-0.0.3-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 673.9 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 561.9 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 673.9 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 562.3 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 748.6 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 663.1 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 741.0 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.6 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 737.2 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-rrf` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 652.9 KiB | [postgresql-15-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-15-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rrf_14` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 840.8 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rrf_14-0.0.3-3PIGSTY.el8.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 753.3 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rrf_14-0.0.3-3PIGSTY.el8.aarch64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 849.4 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rrf_14-0.0.3-3PIGSTY.el9.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 798.2 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rrf_14-0.0.3-3PIGSTY.el9.aarch64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 849.2 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rrf_14-0.0.3-3PIGSTY.el10.x86_64.rpm) |
| `pg_rrf_14` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 785.5 KiB | [pg_rrf_14-0.0.3-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rrf_14-0.0.3-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 672.1 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 560.5 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 671.9 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 560.8 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 745.8 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.1 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 739.8 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 652.5 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 735.1 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-rrf` | `0.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 650.8 KiB | [postgresql-14-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rrf/postgresql-14-pg-rrf_0.0.3-3PIGSTY~resolute_arm64.deb) |

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

pig install pg_rrf -v 18;   # install for PG 18
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
> Sources: [README](https://github.com/yuiseki/pg_rrf/blob/main/README.md), [v0.0.3 release](https://github.com/yuiseki/pg_rrf/releases/tag/v0.0.3)

`pg_rrf` provides Reciprocal Rank Fusion functions for hybrid search score fusion.
It is focused on combining ranked candidate lists without hand-written `FULL OUTER JOIN` / `COALESCE` plumbing.

### Core Functions

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

The `v0.0.3` release explicitly adds `rrfn` while keeping `rrf` and `rrf3` as compatibility wrappers. The README documents the same score behavior:

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

### Notes

The README targets PostgreSQL 14-17 and documents Docker-based build and test flows. The extension surface remains intentionally small: score helpers plus `rrf_fuse` for the common two-list hybrid-search pattern.
