---
title: "hashlib"
linkTitle: "hashlib"
description: "Stable hash functions for Postgres"
weight: 4400
categories: ["UTIL"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_hashlib**](https://github.com/markokr/pghashlib) : Stable hash functions for Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4400** | {{< badge content="hashlib" link="https://github.com/markokr/pghashlib" >}} | {{< ext "hashlib" "pg_hashlib" >}} | `1.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgcrypto" >}} {{< ext "xxhash" >}} {{< ext "md5hash" >}} {{< ext "hashtypes" >}} {{< ext "pgsodium" >}} {{< ext "shacrypt" >}} {{< ext "cryptint" >}} {{< ext "pguecc" >}} |

> [!Note] build-deps: python3-docutils


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_hashlib` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "19" "pg_hashlib_19" "green" >}} {{< bg "18" "pg_hashlib_18" "green" >}} {{< bg "17" "pg_hashlib_17" "green" >}} {{< bg "16" "pg_hashlib_16" "green" >}} {{< bg "15" "pg_hashlib_15" "green" >}} | `pg_hashlib_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "19" "postgresql-19-pg-hashlib" "green" >}} {{< bg "18" "postgresql-18-pg-hashlib" "green" >}} {{< bg "17" "postgresql-17-pg-hashlib" "green" >}} {{< bg "16" "postgresql-16-pg-hashlib" "green" >}} {{< bg "15" "postgresql-15-pg-hashlib" "green" >}} | `postgresql-$v-pg-hashlib` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_hashlib_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_hashlib_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_hashlib_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-19-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-hashlib : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-hashlib : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashlib_19` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.5 KiB | [pg_hashlib_19-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_19-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_hashlib_19` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.8 KiB | [pg_hashlib_19-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_19-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_hashlib_19` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.9 KiB | [pg_hashlib_19-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashlib_19-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_hashlib_19` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [pg_hashlib_19-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashlib_19-1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-pg-hashlib` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 41.9 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.4 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.0 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.5 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.6 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.9 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.8 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.5 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.1 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-pg-hashlib` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 43.9 KiB | [postgresql-19-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-19-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashlib_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.4 KiB | [pg_hashlib_18-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_18-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_hashlib_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.4 KiB | [pg_hashlib_18-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_18-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_hashlib_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.4 KiB | [pg_hashlib_18-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_18-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_hashlib_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.8 KiB | [pg_hashlib_18-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_18-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_hashlib_18` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.8 KiB | [pg_hashlib_18-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashlib_18-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_hashlib_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [pg_hashlib_18-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashlib_18-1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-hashlib` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 41.8 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.2 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 41.9 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.3 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.6 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.0 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.7 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.3 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 43.8 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-hashlib` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 43.7 KiB | [postgresql-18-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-18-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashlib_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.4 KiB | [pg_hashlib_17-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_17-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_hashlib_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.4 KiB | [pg_hashlib_17-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_17-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_hashlib_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.4 KiB | [pg_hashlib_17-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_17-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_hashlib_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.8 KiB | [pg_hashlib_17-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_17-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_hashlib_17` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.8 KiB | [pg_hashlib_17-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashlib_17-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_hashlib_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [pg_hashlib_17-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashlib_17-1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-hashlib` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.0 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.1 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.0 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.3 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.5 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.6 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.8 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.3 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 43.8 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-hashlib` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 43.7 KiB | [postgresql-17-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-17-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashlib_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.4 KiB | [pg_hashlib_16-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_16-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_hashlib_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.4 KiB | [pg_hashlib_16-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_16-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_hashlib_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.4 KiB | [pg_hashlib_16-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_16-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_hashlib_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.8 KiB | [pg_hashlib_16-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_16-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_hashlib_16` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.8 KiB | [pg_hashlib_16-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashlib_16-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_hashlib_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [pg_hashlib_16-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashlib_16-1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-hashlib` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.0 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.2 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.1 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.3 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.5 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.6 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.9 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.3 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 43.8 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-hashlib` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 43.7 KiB | [postgresql-16-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-16-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashlib_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.6 KiB | [pg_hashlib_15-1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashlib_15-1.1-1PGSTY.el8.x86_64.rpm) |
| `pg_hashlib_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.6 KiB | [pg_hashlib_15-1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashlib_15-1.1-1PGSTY.el8.aarch64.rpm) |
| `pg_hashlib_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.6 KiB | [pg_hashlib_15-1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashlib_15-1.1-1PGSTY.el9.x86_64.rpm) |
| `pg_hashlib_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 55.1 KiB | [pg_hashlib_15-1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashlib_15-1.1-1PGSTY.el9.aarch64.rpm) |
| `pg_hashlib_15` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.1 KiB | [pg_hashlib_15-1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashlib_15-1.1-1PGSTY.el10.x86_64.rpm) |
| `pg_hashlib_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 55.2 KiB | [pg_hashlib_15-1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashlib_15-1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-hashlib` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 42.0 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.0 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 42.1 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.2 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.9 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 46.2 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 43.9 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.3 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.1 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-hashlib` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.4 KiB | [postgresql-15-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-hashlib/postgresql-15-pg-hashlib_1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/markokr/pghashlib" title="Repository" icon="github" subtitle="github.com/markokr/pghashlib" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_hashlib-1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_hashlib;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_hashlib;		# install via package name, for the active PG version
pig install hashlib;		# install by extension name, for the current active PG version

pig install hashlib -v 19;   # install for PG 19
pig install hashlib -v 18;   # install for PG 18
pig install hashlib -v 17;   # install for PG 17
pig install hashlib -v 16;   # install for PG 16
pig install hashlib -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION hashlib;
```




## Usage

> [hashlib: Stable hash functions library for PostgreSQL](https://github.com/markokr/pghashlib)

Provides stable hash functions whose implementations do not change across PostgreSQL versions.

### String Hashing (32-bit)

```sql
SELECT hash_string('hello', 'crc32');
SELECT hash_string('hello', 'murmur3');
```

With optional initial value:

```sql
SELECT hash_string('hello', 'crc32', 42);
```

### String Hashing (64-bit)

```sql
SELECT hash64_string('hello', 'city64');
SELECT hash64_string('hello', 'siphash24');
SELECT hash64_string('hello', 'lookup3');
```

### String Hashing (128-bit)

```sql
SELECT hash128_string('hello', 'md5');
SELECT hash128_string('hello', 'city128');
SELECT hash128_string('hello', 'spooky');
```

### Integer Hashing

```sql
SELECT hash_int4(42);        -- 32-bit hash of 32-bit integer
SELECT hash_int8(42::bigint); -- 64-bit hash of 64-bit integer
```

### Available Algorithms

| Algorithm | CPU-indep | Bits | Description |
|-----------|-----------|------|-------------|
| `crc32` | yes | 32 | CRC32 |
| `murmur3` | no | 32 | MurmurHash v3 |
| `md5` | yes | 128 | MD5 |
| `city64` | no | 64 | CityHash64 |
| `city128` | no | 128 | CityHash128 |
| `siphash24` | yes | 64 | SipHash-2-4 |
| `spooky` | no | 128 | SpookyHash |
| `lookup2` | no | 64 | Jenkins lookup2 |
| `lookup3` | no | 64 | Jenkins lookup3 CPU-native |
| `lookup3be` | yes | 64 | Jenkins lookup3 big-endian |
| `lookup3le` | yes | 64 | Jenkins lookup3 little-endian |
| `pgsql84` | no | 64 | Hacked lookup3 in Postgres 8.4+ |
{.wide .expand}

Integer algorithms: `wang32`, `wang32mult`, `jenkins` (32-bit); `wang64`, `wang64to32` (64-bit). All are reversible (1:1 mapping), useful for creating random sort orders over unique IDs.
