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
| **2890** | {{< badge content="pg_local_cache" link="https://github.com/profundium/pg_local_cache" >}} | {{< ext "pg_local_cache" >}} | `1.3.0` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `local_cache` |

> [!Note] Requires shared_preload_libraries=pg_local_cache and a restart; CREATE EXTENSION requires superuser; v1.3.0 supports PostgreSQL 14-18, one configured database, and one writable primary.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_local_cache` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "18" "pg_local_cache_18" "green" >}} {{< bg "17" "pg_local_cache_17" "green" >}} {{< bg "16" "pg_local_cache_16" "green" >}} {{< bg "15" "pg_local_cache_15" "green" >}} {{< bg "14" "pg_local_cache_14" "green" >}} | `pg_local_cache_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "18" "postgresql-18-pg-local-cache" "green" >}} {{< bg "17" "postgresql-17-pg-local-cache" "green" >}} {{< bg "16" "postgresql-16-pg-local-cache" "green" >}} {{< bg "15" "postgresql-15-pg-local-cache" "green" >}} {{< bg "14" "postgresql-14-pg-local-cache" "green" >}} | `postgresql-$v-pg-local-cache` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "pg_local_cache_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-18-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-17-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pg-local-cache : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-14-pg-local-cache : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_18` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 91.8 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_18-1.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_local_cache_18` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 88.7 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_18-1.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_local_cache_18` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.7 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_18-1.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_local_cache_18` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 87.4 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_18-1.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_local_cache_18` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 89.4 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_18-1.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_local_cache_18` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 88.2 KiB | [pg_local_cache_18-1.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_18-1.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 211.8 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 205.8 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.0 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 206.9 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 230.8 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 226.4 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.1 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 219.5 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 218.8 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-local-cache` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 216.1 KiB | [postgresql-18-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-18-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_17` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 91.8 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_17-1.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_local_cache_17` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 88.7 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_17-1.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_local_cache_17` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.7 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_17-1.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_local_cache_17` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 87.4 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_17-1.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_local_cache_17` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 89.5 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_17-1.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_local_cache_17` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 88.1 KiB | [pg_local_cache_17-1.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_17-1.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 211.7 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 205.8 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 212.5 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 206.6 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 251.8 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 247.8 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.3 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 219.5 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 218.8 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-local-cache` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 215.2 KiB | [postgresql-17-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-17-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_16` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 91.8 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_16-1.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_local_cache_16` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 88.7 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_16-1.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_local_cache_16` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 88.7 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_16-1.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_local_cache_16` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 87.5 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_16-1.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_local_cache_16` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 89.5 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_16-1.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_local_cache_16` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 88.2 KiB | [pg_local_cache_16-1.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_16-1.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 211.6 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 206.1 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 212.5 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 206.6 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 251.0 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 246.4 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 220.9 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 219.5 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 218.8 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-local-cache` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 215.4 KiB | [postgresql-16-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-16-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_15` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 93.3 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_15-1.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_local_cache_15` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 90.4 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_15-1.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_local_cache_15` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 90.8 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_15-1.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_local_cache_15` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 90.5 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_15-1.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_local_cache_15` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.9 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_15-1.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_local_cache_15` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.7 KiB | [pg_local_cache_15-1.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_15-1.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 213.2 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.0 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.5 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.6 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 252.0 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 249.1 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 222.7 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 221.5 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 220.1 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-local-cache` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 217.3 KiB | [postgresql-15-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-15-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_local_cache_14` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 93.3 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_local_cache_14-1.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_local_cache_14` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 90.6 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_local_cache_14-1.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_local_cache_14` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 90.7 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_local_cache_14-1.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_local_cache_14` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.5 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_local_cache_14-1.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_local_cache_14` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.8 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_local_cache_14-1.3.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_local_cache_14` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.6 KiB | [pg_local_cache_14-1.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_local_cache_14-1.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 212.8 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.8 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.3 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 208.5 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 248.2 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 246.7 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 222.5 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 222.2 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 219.8 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-local-cache` | `1.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 218.4 KiB | [postgresql-14-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-local-cache/postgresql-14-pg-local-cache_1.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/profundium/pg_local_cache" title="Repository" icon="github" subtitle="github.com/profundium/pg_local_cache" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_local_cache-1.3.0.tar.gz" />}}
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

- [pg_local_cache 1.3.0 on PGXN](https://pgxn.org/dist/pg_local_cache/1.3.0/)
- [pg_local_cache v1.3.0 README](https://github.com/profundium/pg_local_cache/blob/v1.3.0/README.md)
- [pg_local_cache v1.3.0 control file](https://github.com/profundium/pg_local_cache/blob/v1.3.0/pg_local_cache.control)
- [pg_local_cache 1.3.0 extension SQL](https://github.com/profundium/pg_local_cache/blob/v1.3.0/sql/pg_local_cache--1.3.0.sql)
- [Technical reference](https://github.com/profundium/pg_local_cache/blob/v1.3.0/docs/TECHNICAL.md)
- [Existing-server installation guide](https://github.com/profundium/pg_local_cache/blob/v1.3.0/docs/INSTALL_EXISTING.md)
- [pg_local_cache v1.3.0 release](https://github.com/profundium/pg_local_cache/releases/tag/v1.3.0)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_local_cache)

`pg_local_cache` 1.3.0 is a transaction-aware, in-process cache for repeated PostgreSQL primary-key reads. It keeps bounded whole-row entries in shared memory and can transparently accelerate eligible ordinary `SELECT` statements while retaining the original PostgreSQL primary-key plan as the authoritative fallback. Use it for a hot working set on one writable primary; it is not a general query-result cache, a durability layer, or a distributed Redis/Valkey replacement.

### Core Workflow

The library must be loaded at postmaster startup. This SQL-only configuration disables the optional RESP listener and serves one application database:

```conf
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.port = 0
pg_local_cache.cache_entries = 16384
pg_local_cache.memory_budget_mb = 384
```

Add `pg_local_cache` to any existing comma-separated preload list instead of replacing other libraries, validate the configuration, and perform a controlled PostgreSQL restart. The control file fixes the extension in schema `local_cache`, sets `superuser=true`, and is not relocatable, so a superuser must create it in each database where it is used:

```sql
CREATE EXTENSION pg_local_cache;
```

Create an eligible permanent table, then attach it. `attach_table` takes a `ShareRowExclusiveLock`, records the complete primary key in `local_cache.mapping`, installs extension-owned invalidation triggers, and publishes the mapping to shared memory. Use a bounded lock timeout on a live system:

```sql
CREATE TABLE public.items (
    id bigint PRIMARY KEY,
    value text NOT NULL,
    enabled boolean NOT NULL DEFAULT true,
    metadata jsonb
);

INSERT INTO public.items VALUES
    (1, 'hello', true, '{"source":"postgres"}');

BEGIN;
SET LOCAL lock_timeout = '2s';
SELECT local_cache.attach_table('public.items'::regclass);
COMMIT;
```

The default `p_writable=false` disables RESP `SET` and `DEL`; it does not prevent normal PostgreSQL DML. Applications keep using their existing PostgreSQL connection, row types, and SQL:

```sql
SELECT * FROM public.items WHERE id = $1::bigint;

SELECT value, metadata
FROM public.items
WHERE id = ANY($1::bigint[]);

EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM public.items WHERE id = 1;
```

An eligible plan appears as `Custom Scan (pg_local_cache_sql)`. A cache miss or any unsafe or unsupported condition executes the retained primary-key index plan; PostgreSQL remains the source of truth.

### Explicit JSON APIs

Ordinary SQL is the canonical tuple-returning interface. Callers that deliberately want a cache-shaped JSON API can use these `SECURITY INVOKER` functions:

```sql
SELECT local_cache.get('public.items'::regclass, 1::bigint);

SELECT local_cache.mget(
    'public.items'::regclass,
    ARRAY[1, 7, 1]::bigint[]
);
```

`get(regclass, anyelement)` returns complete-row JSON as `text`; `mget(regclass, anyarray)` returns a `text[]` aligned with its input, preserving duplicate and `NULL` positions. For a composite or heterogeneous primary key, call `get(regclass, text[])` with components in the key order recorded by `attach_table`. Explicit API callers need `USAGE` on schema `local_cache`, `EXECUTE` on the chosen overload, and normal `SELECT` privilege on the source table.

### Important Objects and Controls

- `local_cache.attach_table(regclass, boolean, text)` validates and registers a table. Set `p_writable=true` only when the optional RESP worker should be allowed to write the source relation; `p_namespace` overrides the generated mapping name.
- `local_cache.detach_table(regclass)` removes the mapping, managed triggers, shared entry state, and direct worker-role privileges. It returns `false` when the relation was not attached.
- `local_cache.reconcile_table(regclass)` and `local_cache.reconcile_all()` revalidate relation shape, primary keys, trigger provenance, and worker grants after controlled DDL or privilege changes.
- `local_cache.mapping` is the extension-owned mapping registry and is included in extension configuration dumps. Do not edit it as a substitute for the administrative functions.
- `local_cache.metrics()` returns typed counters and memory/worker gauges, `local_cache.health()` returns a compact JSON readiness assessment, and `local_cache.stats()` returns detailed JSON diagnostics. These and the administrative functions are revoked from `PUBLIC`; grant them only to designated deploy or monitoring roles.
- `local_cache.invalidate(namespace)` invalidates one mapping namespace and returns the affected-entry count. Normal DML, `TRUNCATE`, and relevant DDL use automatic transactional invalidation.

Key settings are:

| Setting | Default | Effect |
|---|---:|---|
| `pg_local_cache.port` | `6380` | RESP2 port; set `0` for SQL-only mode. |
| `pg_local_cache.database` | `postgres` | The one database served by this extension instance. |
| `pg_local_cache.cache_entries` | `16384` | Shared row-cache entry capacity. |
| `pg_local_cache.relation_states` | `1024` | Shared relation-version state capacity. |
| `pg_local_cache.memory_budget_mb` | `384` | Startup budget for deterministic extension allocations. |
| `pg_local_cache.max_dirty_keys` | `4096` | Per-transaction key bound before invalidation widens to the relation. |
| `pg_local_cache.sql_cache` | `on` | `USERSET` switch for the ordinary-SQL fast path; no restart is required. |

Except for `pg_local_cache.sql_cache`, the documented GUCs are postmaster settings. The memory budget covers the extension's deterministic shared hashes and optional RESP buffers, not `shared_buffers`, backend memory, the operating system, or other services.

### Fast-Path and Consistency Boundaries

The transparent path is deliberately narrow. It requires `READ COMMITTED`, one attached base table, direct column projections, and equality predicates for every primary-key column. A single-column primary key also supports bounded `IN` and `= ANY(array)` queries. Joins, CTEs, subqueries, aggregates, grouping, ordering, row locks, extra predicates, recovery, parallel execution, `REPEATABLE READ`, and `SERIALIZABLE` use normal PostgreSQL plans. Scalar lookups may use no `LIMIT` or constant `LIMIT 1`; batch lookups may not use `LIMIT`.

For `IN`/`ANY`, the executor deduplicates at most 1,024 non-null keys and copies at most 16 MiB of validated rows into query-local memory. The batch is all-or-nothing: one miss, unsafe snapshot, malformed entry, or budget overflow runs the complete source plan rather than merging cached and source rows.

Source-table writes remain ordinary PostgreSQL transactions. Managed triggers collect changed keys, and the pre-commit callback publishes invalidation fences before the transaction becomes visible. A rollback never publishes uncommitted row data. After the current transaction writes an attached relation, subsequent reads in that transaction bypass the cache to preserve read-your-own-write behavior. `PREPARE TRANSACTION` is rejected after such a write.

Entries have no TTL. They remain until invalidation, eviction, replacement, corruption detection, or an MVCC safety check retires them. Encoded cache values are limited to 8 KiB; a wider row simply uses PostgreSQL instead of becoming an entry.

### Table and Deployment Requirements

Version 1.3.0 supports PostgreSQL 14–18, one configured database, and one writable primary. Upstream's own published binary and existing-server instructions cover Linux amd64 with glibc or musl; the current Pigsty package matrix separately includes validated x86_64 and aarch64 builds. Treat those as different evidence layers and verify the exact package platform before installation. Attached relations must be permanent, non-partitioned tables with no inheritance or row-level security and with an immediate, non-partial B-tree primary key. Supported key columns are `int2`, `int4`, `int8`, `text`, `varchar`, `bpchar`, and `uuid`; composite keys may contain 1–16 columns. Temporary or unlogged tables, views, partitioned tables, expression or partial primary keys, nondeterministic key collations, and non-default primary-key operator classes are rejected.

At most 128 mappings are published per instance. Dropping a table forgets its mapping; recreating a table with the same name does not reattach it. The cache is not served on standbys and provides no multi-primary coordination, TTL, clustering, Pub/Sub, or general range/join/aggregate caching.

### Version 1.3.0 Upgrade

Version 1.3.0 changes the shared library, packaging, or documentation; its SQL objects are unchanged from 1.2.1. Because the library is loaded at postmaster startup, install the matching files, perform a controlled restart, then record the extension version:

```sql
ALTER EXTENSION pg_local_cache UPDATE TO '1.3.0';
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_local_cache';
```

After restart, check `local_cache.health()`, `local_cache.metrics()`, attached mappings, and an `EXPLAIN (ANALYZE, COSTS OFF)` fast-path query before returning traffic. Do not infer runtime readiness from the extension version alone.

### Optional RESP2 Security Boundary

RESP mode exposes whole-row `GET`, `SET`, and `DEL` through a limited RESP2 protocol, but it uses one shared token and one `LOGIN NOSUPERUSER NOINHERIT` worker role for every accepted mapping. It has no TLS and no per-client PostgreSQL identity or ACL context. Keep `pg_local_cache.port=0` unless this interface is required. If enabled, retain the default loopback bind or place remote access behind network isolation and authenticated TLS, store the token in a PostgreSQL OS-user-owned mode `0400` or `0600` file through `pg_local_cache.auth_token_file`, and never treat a lost write reply as proof that the PostgreSQL transaction did not commit.
