---
title: "xxhash"
linkTitle: "xxhash"
description: "xxhash functions for PostgreSQL"
weight: 4430
categories: ["UTIL"]
width: full
---

[**pg_xxhash**](https://github.com/hatarist/pg_xxhash) : xxhash functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4430** | {{< badge content="xxhash" link="https://github.com/hatarist/pg_xxhash" >}} | {{< ext "xxhash" "pg_xxhash" >}} | `0.0.1` | {{< category "UTIL" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "shacrypt" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_xxhash` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_xxhash_18" "green" >}} {{< bg "17" "pg_xxhash_17" "green" >}} {{< bg "16" "pg_xxhash_16" "green" >}} {{< bg "15" "pg_xxhash_15" "green" >}} {{< bg "14" "pg_xxhash_14" "green" >}} {{< bg "13" "pg_xxhash_13" "green" >}} | `pg_xxhash_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-xxhash" "green" >}} {{< bg "17" "postgresql-17-pg-xxhash" "green" >}} {{< bg "16" "postgresql-16-pg-xxhash" "green" >}} {{< bg "15" "postgresql-15-pg-xxhash" "green" >}} {{< bg "14" "postgresql-14-pg-xxhash" "green" >}} {{< bg "13" "postgresql-13-pg-xxhash" "green" >}} | `postgresql-$v-pg-xxhash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_xxhash_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-xxhash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-xxhash : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.3 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.0 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.2 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.6 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.4 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 76.6 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.9 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.4 KiB | [postgresql-18-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-18-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.2 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.0 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.2 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.6 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.9 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 78.7 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.8 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.4 KiB | [postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-17-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.2 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.8 KiB | [pg_xxhash_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.2 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.0 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.2 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.5 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 76.9 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 78.7 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.8 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 67.4 KiB | [postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-16-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.7 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.6 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.1 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.5 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_xxhash_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.0 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.8 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.9 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.3 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 87.0 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 89.7 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 73.4 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 78.4 KiB | [postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-15-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.7 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.6 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.1 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.5 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_xxhash_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.9 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 74.0 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.9 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.5 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 87.1 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 89.5 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 73.4 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 78.6 KiB | [postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-14-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xxhash_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.7 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xxhash_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_xxhash_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.6 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xxhash_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_xxhash_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.0 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xxhash_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_xxhash_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.5 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xxhash_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_xxhash_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xxhash_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_xxhash_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_xxhash_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xxhash_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.9 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 74.1 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.8 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 74.6 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 87.1 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 89.5 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 73.4 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-xxhash` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 78.6 KiB | [postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xxhash/postgresql-13-pg-xxhash_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hatarist/pg_xxhash" title="Repository" icon="github" subtitle="github.com/hatarist/pg_xxhash" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_xxhash-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_xxhash;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_xxhash;		# install via package name, for the active PG version
pig install xxhash;		# install by extension name, for the current active PG version

pig install xxhash -v 18;   # install for PG 18
pig install xxhash -v 17;   # install for PG 17
pig install xxhash -v 16;   # install for PG 16
pig install xxhash -v 15;   # install for PG 15
pig install xxhash -v 14;   # install for PG 14
pig install xxhash -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION xxhash;
```
