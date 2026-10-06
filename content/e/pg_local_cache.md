---
title: "pg_local_cache"
linkTitle: "pg_local_cache"
description: "Transaction-aware shared-memory cache for ordinary PostgreSQL primary-key reads"
weight: 2890
categories: ["FEAT"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_local_cache**](https://github.com/profundium/pg_local_cache) : Transaction-aware shared-memory cache for ordinary PostgreSQL primary-key reads


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2890** | {{< badge content="pg_local_cache" link="https://github.com/profundium/pg_local_cache" >}} | {{< ext "pg_local_cache" >}} | `2.0.4` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `local_cache` |

> [!Note] PG14-18; requires preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_local_cache` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "pg_local_cache_18" "green" >}} {{< bg "17" "pg_local_cache_17" "green" >}} {{< bg "16" "pg_local_cache_16" "green" >}} {{< bg "15" "pg_local_cache_15" "green" >}} {{< bg "14" "pg_local_cache_14" "green" >}} | `pg_local_cache_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "postgresql-18-pg-local-cache" "green" >}} {{< bg "17" "postgresql-17-pg-local-cache" "green" >}} {{< bg "16" "postgresql-16-pg-local-cache" "green" >}} {{< bg "15" "postgresql-15-pg-local-cache" "green" >}} {{< bg "14" "postgresql-14-pg-local-cache" "green" >}} | `postgresql-$v-pg-local-cache` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_18` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 187.9 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_18-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_local_cache_18` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 182.9 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_18-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_local_cache_18` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 189.3 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_18-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_local_cache_18` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 185.8 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_18-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_local_cache_18` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 190.2 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_18-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_local_cache_18` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 185.9 KiB | [pg_local_cache_18-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_18-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 173.2 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 168.1 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 173.4 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 168.4 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 187.0 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 182.8 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 180.7 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 177.7 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 178.2 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-local-cache` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 175.1 KiB | [postgresql-18-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_17` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 187.9 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_17-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_local_cache_17` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 182.8 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_17-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_local_cache_17` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 188.8 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_17-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_local_cache_17` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 185.6 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_17-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_local_cache_17` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 189.6 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_17-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_local_cache_17` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 185.7 KiB | [pg_local_cache_17-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_17-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 173.0 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 167.9 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 173.5 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 168.1 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 204.0 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 200.4 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 180.3 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 177.5 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 177.9 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-local-cache` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 175.0 KiB | [postgresql-17-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 187.9 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_16-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_local_cache_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 182.8 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_16-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_local_cache_16` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 188.8 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_16-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_local_cache_16` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 185.6 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_16-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_local_cache_16` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 189.6 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_16-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_local_cache_16` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 185.7 KiB | [pg_local_cache_16-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_16-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 173.0 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 167.8 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 173.3 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 168.1 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 203.0 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 199.3 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 180.6 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 177.5 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 178.0 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-local-cache` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 174.8 KiB | [postgresql-16-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 188.5 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_15-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_local_cache_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.8 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_15-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_local_cache_15` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 190.8 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_15-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_local_cache_15` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 188.3 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_15-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_local_cache_15` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 191.4 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_15-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_local_cache_15` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 188.4 KiB | [pg_local_cache_15-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_15-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 174.3 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 169.0 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 174.8 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 169.1 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 204.6 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 201.3 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 182.7 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 179.9 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 180.1 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-local-cache` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 176.9 KiB | [postgresql-15-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_14` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 188.3 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_14-2.0.4-1PGSTY.el8.x86_64.rpm) |
| `pg_local_cache_14` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 184.7 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_14-2.0.4-1PGSTY.el8.aarch64.rpm) |
| `pg_local_cache_14` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 190.6 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_14-2.0.4-1PGSTY.el9.x86_64.rpm) |
| `pg_local_cache_14` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 189.9 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_14-2.0.4-1PGSTY.el9.aarch64.rpm) |
| `pg_local_cache_14` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 191.5 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_14-2.0.4-1PGSTY.el10.x86_64.rpm) |
| `pg_local_cache_14` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 189.9 KiB | [pg_local_cache_14-2.0.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_14-2.0.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 174.1 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 169.6 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 174.4 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 170.2 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 200.8 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 199.0 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 182.3 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 181.0 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 179.6 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-local-cache` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 178.6 KiB | [postgresql-14-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_2.0.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/profundium/pg_local_cache" title="Repository" icon="github" subtitle="github.com/profundium/pg_local_cache" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_local_cache-2.0.4.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_local_cache;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_local_cache;		# install via package name, for the active PG version

pig install pg_local_cache -v 18;   # install for PG 18
pig install pg_local_cache -v 17;   # install for PG 17
pig install pg_local_cache -v 16;   # install for PG 16
pig install pg_local_cache -v 15;   # install for PG 15
pig install pg_local_cache -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_local_cache';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_local_cache;
```

## Usage

Sources:

- [README.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/README.md)
- [pg_local_cache.control](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/pg_local_cache.control)
- [sql/pg_local_cache--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--3.0.0.sql)
- [sql/pg_local_cache--2.0.4--3.0.0.sql](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/sql/pg_local_cache--2.0.4--3.0.0.sql)
- [docs/UPGRADING.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/docs/UPGRADING.md)
- [CHANGELOG.md](https://github.com/profundium/pg_local_cache/blob/940f796de41bf4ad22de3f8efc32dcf6d6b791b9/CHANGELOG.md)

`pg_local_cache` 3.0.0 caches whole rows by their complete primary key in bounded shared memory. Cached reads now use authenticated RESP MGET; SQL `local_cache.mget` is removed and ordinary SELECT queries are not rewritten. PostgreSQL remains the durable source of truth.

### Core Workflow

```ini
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.role = 'local_cache_worker'
pg_local_cache.bind_address = '127.0.0.1'
pg_local_cache.port = 6380
pg_local_cache.auth_token_file = '/secure/path/token'
```

```sql
CREATE EXTENSION pg_local_cache;
CREATE TABLE public.items (id bigint PRIMARY KEY, value text);
INSERT INTO public.items VALUES (42, 'example');
SELECT local_cache.attach_table('public.items'::regclass);
SELECT local_cache.health();
```

After authenticating the RESP connection:

```text
MGET CRUD:app.public.items:{"id":42} CRUD:app.public.items:{"id":7}
```

### Operational Boundaries

Append the library to existing preload entries, provision a dedicated worker role and its documented metadata/table grants, secure the token file, and restart. A superuser creates the extension in fixed schema `local_cache`. Attach only supported permanent primary-key tables. RESP MGET preserves order and duplicate keys and returns null for missing rows. Workers use the configured database role, not the client’s SQL privileges, transaction or snapshot. Keep joins, projections, row locks and session-sensitive reads in ordinary SQL.

`local_cache.attach_table` installs invalidation triggers; `local_cache.detach_table` removes a mapping; `local_cache.reconcile_table` revalidates it after DDL or privilege changes. `local_cache.health`, `local_cache.stats` and `local_cache.metrics` report readiness and resource use. Normal PostgreSQL writes invalidate affected cache entries. RLS, partitioned and inherited tables are outside the documented workload. There is no TTL or distributed cache coordination.

Supports PostgreSQL 14–18 on one writable primary. Non-loopback listeners require `pg_local_cache.tls` or explicit `pg_local_cache.allow_plaintext_network` opt-in. TLS uses separate certificate/key settings; a CA setting enables mutual TLS. The reloadable `pg_local_cache.enabled` switch disables caching.

Before upgrading from 2.x, move SQL MGET callers to the RESP authorization model. Install matching 3.0.0 files, configure listener security, restart, confirm the library version and listener readiness, then run `ALTER EXTENSION pg_local_cache UPDATE` in each affected database. Dependencies on removed functions or the replaced metrics return type can block migration; the script deliberately avoids CASCADE. No downgrade script is supplied.
