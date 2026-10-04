---
title: "acdat"
linkTitle: "acdat"
description: "Compiled Aho-Corasick double-array machines for exact multi-pattern matching and replacement in PostgreSQL"
weight: 2260
categories: ["FTS"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**acdat**](https://github.com/pgsty/acdat) : Compiled Aho-Corasick double-array machines for exact multi-pattern matching and replacement in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2260** | {{< badge content="acdat" link="https://github.com/pgsty/acdat" >}} | {{< ext "acdat" >}} | `0.1.1` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `acdat` |
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_bigm" >}} {{< ext "pgroonga" >}} {{< ext "pg_search" >}} |

> [!Note] 0.1.1 preserves the 0.1.0 SQL API and format-major-1 machine compatibility; fixed acdat schema; no preload required.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `acdat` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "acdat_18" "green" >}} {{< bg "17" "acdat_17" "green" >}} {{< bg "16" "acdat_16" "green" >}} {{< bg "15" "acdat_15" "green" >}} {{< bg "14" "acdat_14" "green" >}} | `acdat_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-acdat" "green" >}} {{< bg "17" "postgresql-17-acdat" "green" >}} {{< bg "16" "postgresql-16-acdat" "green" >}} {{< bg "15" "postgresql-15-acdat" "green" >}} {{< bg "14" "postgresql-14-acdat" "green" >}} | `postgresql-$v-acdat` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "acdat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "acdat_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-acdat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-acdat : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acdat_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 118.5 KiB | [acdat_18-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acdat_18-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `acdat_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 116.9 KiB | [acdat_18-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acdat_18-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `acdat_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.7 KiB | [acdat_18-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acdat_18-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `acdat_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 117.5 KiB | [acdat_18-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acdat_18-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `acdat_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 121.2 KiB | [acdat_18-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acdat_18-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `acdat_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 120.0 KiB | [acdat_18-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acdat_18-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-acdat` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 109.9 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 107.8 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.1 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.1 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.2 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 114.9 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.7 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.8 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.7 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-acdat` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 111.8 KiB | [postgresql-18-acdat_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-18-acdat_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acdat_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 118.4 KiB | [acdat_17-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acdat_17-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `acdat_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 116.7 KiB | [acdat_17-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acdat_17-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `acdat_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.5 KiB | [acdat_17-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acdat_17-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `acdat_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 117.3 KiB | [acdat_17-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acdat_17-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `acdat_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 121.0 KiB | [acdat_17-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acdat_17-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `acdat_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 119.8 KiB | [acdat_17-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acdat_17-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-acdat` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 110.0 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 107.7 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.3 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.1 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 122.3 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 121.0 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.7 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.8 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.7 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-acdat` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 111.8 KiB | [postgresql-17-acdat_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-17-acdat_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acdat_16` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 118.5 KiB | [acdat_16-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acdat_16-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `acdat_16` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 116.7 KiB | [acdat_16-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acdat_16-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `acdat_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.5 KiB | [acdat_16-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acdat_16-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `acdat_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 117.2 KiB | [acdat_16-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acdat_16-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `acdat_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 121.0 KiB | [acdat_16-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acdat_16-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `acdat_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 119.7 KiB | [acdat_16-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acdat_16-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-acdat` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 110.0 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 107.7 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.3 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.2 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 122.2 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 120.9 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.8 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.8 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.7 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-acdat` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 111.8 KiB | [postgresql-16-acdat_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-16-acdat_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acdat_15` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 119.0 KiB | [acdat_15-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acdat_15-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `acdat_15` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.2 KiB | [acdat_15-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acdat_15-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `acdat_15` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 120.6 KiB | [acdat_15-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acdat_15-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `acdat_15` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 119.2 KiB | [acdat_15-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acdat_15-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `acdat_15` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.8 KiB | [acdat_15-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acdat_15-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `acdat_15` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.5 KiB | [acdat_15-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acdat_15-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-acdat` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 110.9 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.1 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.6 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.7 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 123.9 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 122.7 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.5 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.6 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 114.2 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-acdat` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 113.5 KiB | [postgresql-15-acdat_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-15-acdat_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acdat_14` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 119.0 KiB | [acdat_14-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acdat_14-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `acdat_14` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.1 KiB | [acdat_14-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acdat_14-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `acdat_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 120.6 KiB | [acdat_14-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acdat_14-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `acdat_14` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 119.1 KiB | [acdat_14-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acdat_14-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `acdat_14` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.8 KiB | [acdat_14-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acdat_14-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `acdat_14` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.5 KiB | [acdat_14-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acdat_14-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-acdat` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 110.9 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.1 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.9 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.7 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 123.9 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 122.7 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.5 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.6 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 114.3 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-acdat` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 113.6 KiB | [postgresql-14-acdat_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/acdat/postgresql-14-acdat_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsty/acdat" title="Repository" icon="github" subtitle="github.com/pgsty/acdat" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="acdat-0.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg acdat;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install acdat;		# install via package name, for the active PG version

pig install acdat -v 18;   # install for PG 18
pig install acdat -v 17;   # install for PG 17
pig install acdat -v 16;   # install for PG 16
pig install acdat -v 15;   # install for PG 15
pig install acdat -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION acdat;
```

## Usage

Sources:

- [Official release v0.1.1](https://github.com/pgsty/acdat/releases/tag/v0.1.1)
- [Official README v0.1.1](https://github.com/pgsty/acdat/blob/v0.1.1/README.md)
- [Extension control file](https://github.com/pgsty/acdat/blob/v0.1.1/acdat.control)
- [Versioned installation SQL](https://github.com/pgsty/acdat/blob/v0.1.1/sql/acdat--0.1.1.sql)
- [Official usage guide](https://github.com/pgsty/acdat/blob/v0.1.1/docs/USAGE.md)
- [Runnable SQL demonstration](https://github.com/pgsty/acdat/blob/v0.1.1/examples/demo.sql)

`acdat` 0.1.1 compiles a large dictionary of exact literal patterns into an immutable Aho-Corasick Double-Array machine, then scans each `text` or `bytea` value once for matching or replacement. It is designed for stable, repeatedly used dictionaries such as policy rules, indicators of compromise, entity names, and redaction aliases.

### Core Workflow

Create the extension, compile a dictionary, and reuse the resulting `acdat.machine` value across many inputs:

```sql
CREATE EXTENSION acdat;

WITH machine AS (
    SELECT acdat.compile(
        ARRAY['he', 'she', 'his', 'hers'],
        ARRAY[1, 2, 3, 4]::bigint[]
    ) AS value
)
SELECT acdat.contains('ushers', value) AS matched,
       acdat.info(value)->>'pattern_count' AS patterns
FROM machine;
```

For production dictionaries, the source rules should stay in an application-owned table. The aggregate overload of `acdat.compile()` can build one deterministic machine directly from pattern, ID, replacement, and priority rows; compile once and scan many values.

### Matching and Replacement

`acdat.contains()` stops after the first hit. `acdat.matches()` returns `acdat.hit` rows with the pattern ID, byte and character coordinates, and priority. `acdat.replace()` applies literal, non-recursive replacements:

```sql
WITH machine AS (
    SELECT acdat.compile(
        ARRAY['病毒', '特征码', '病毒特征码'],
        ARRAY[10, 11, 12]::bigint[],
        ARRAY['[VIRUS]', '[SIGNATURE]', '[IOC]'],
        ARRAY[20, 20, 5]::integer[]
    ) AS value
)
SELECT *
FROM acdat.matches('发现病毒特征码', (SELECT value FROM machine), 'all_overlapping');

SELECT acdat.replace(
    'aaa',
    acdat.compile(
        ARRAY['a', 'aa', 'aaa'],
        ARRAY[1, 2, 3]::bigint[],
        ARRAY['[x]', '[yy]', '[zzz]']
    ),
    'leftmost_longest'
);
```

The match policies are `all_overlapping`, `leftmost_longest`, and `leftmost_priority`. Replacement accepts only a non-overlapping policy. Use `acdat.info()` to inspect a compiled machine and the export, validation, import, and fingerprint functions when moving or checking artifacts.

`acdat.matches()` defaults `max_matches` to 10000, and `acdat.replace()` defaults `max_output_bytes` to 268435456. Set tighter limits for untrusted or high-hit inputs so match enumeration and replacement output stay bounded.

### Managed Dictionaries

The optional catalog layer publishes immutable, content-addressed builds and atomically selects one active build. Its control functions use `SECURITY INVOKER` and are not executable by `PUBLIC`:

```sql
WITH machine AS (
    SELECT acdat.compile(pattern, pattern_id)
    FROM app_keyword
    WHERE enabled
), published AS (
    SELECT acdat.publish('moderation', 1, machine) AS build_id
    FROM machine
)
SELECT acdat.activate('moderation', build_id)
FROM published;

SELECT name, version, build_id, machine
FROM acdat.active_machine
WHERE name = 'moderation';
```

Application tables remain the source of truth. Logical dumps include catalog metadata and active machine payloads, but not every historical artifact, so retain the source patterns required to rebuild retired or inactive versions.

### Compatibility and Safety

Version 0.1.1 is tested on PostgreSQL 14 through 18. It needs no preload or server restart, has no external extension dependency, and defines no GUC. The control file fixes the schema to `acdat`, sets `relocatable = false` and `trusted = false`, so `CREATE EXTENSION` requires a superuser.

The 0.1.1 release preserves the 0.1.0 SQL API and format-major-1 machine compatibility and ships the `0.1.0 -> 0.1.1` extension update path. It also adds cancellable compilation, a conservative build-work budget, and a faster materialized scan path without changing the stored machine contract.

ACDAT indexes the pattern dictionary, not the document table: scanning a large existing table still reads its candidate rows. Matching is exact and case-sensitive; the extension does not provide regular expressions, fuzzy matching, tokenization, automatic case folding, Unicode normalization, or a document-side index. The text engine supports UTF-8 and single-byte server encodings, while binary data should use the bytea interface. Materialize `(document_id, pattern_id)` hits into an application table when repeated reverse lookup is required.

The compiled format is self-describing and checksummed, and imported artifacts are validated before use. Inventory dependencies before uninstalling: `DROP EXTENSION acdat` removes managed dictionary state, while adding `CASCADE` can also remove user columns or other objects that depend on `acdat.machine`.
