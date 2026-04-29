---
title: "byteamagic"
linkTitle: "byteamagic"
description: "Detect MIME types and file formats from PostgreSQL bytea values"
weight: 4275
categories: ["UTIL"]
width: full
---

[**pg_byteamagic**](https://github.com/nmandery/pg_byteamagic) : Detect MIME types and file formats from PostgreSQL bytea values


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4275** | {{< badge content="byteamagic" link="https://github.com/nmandery/pg_byteamagic" >}} | {{< ext "byteamagic" "pg_byteamagic" >}} | `0.2.4` | {{< category "UTIL" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Extension name is byteamagic; package name is pg_byteamagic.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_byteamagic` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.4` | {{< bg "18" "pg_byteamagic_18" "green" >}} {{< bg "17" "pg_byteamagic_17" "green" >}} {{< bg "16" "pg_byteamagic_16" "green" >}} {{< bg "15" "pg_byteamagic_15" "green" >}} {{< bg "14" "pg_byteamagic_14" "green" >}} | `pg_byteamagic_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.4` | {{< bg "18" "postgresql-18-pg-byteamagic" "green" >}} {{< bg "17" "postgresql-17-pg-byteamagic" "green" >}} {{< bg "16" "postgresql-16-pg-byteamagic" "green" >}} {{< bg "15" "postgresql-15-pg-byteamagic" "green" >}} {{< bg "14" "postgresql-14-pg-byteamagic" "green" >}} | `postgresql-$v-pg-byteamagic` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "pg_byteamagic_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-18-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-17-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-16-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-15-pg-byteamagic : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.4" "postgresql-14-pg-byteamagic : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-byteamagic : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-byteamagic : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-byteamagic : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_byteamagic_18` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_byteamagic_18-0.2.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_byteamagic_18` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_byteamagic_18-0.2.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_byteamagic_18` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_byteamagic_18-0.2.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_byteamagic_18` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_byteamagic_18-0.2.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_byteamagic_18` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_byteamagic_18-0.2.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_byteamagic_18` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_18-0.2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_byteamagic_18-0.2.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.7 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.7 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.6 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.7 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-byteamagic` | `0.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 KiB | [postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-18-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_byteamagic_17` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_byteamagic_17-0.2.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_byteamagic_17` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_byteamagic_17-0.2.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_byteamagic_17` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_byteamagic_17-0.2.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_byteamagic_17` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_byteamagic_17-0.2.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_byteamagic_17` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_byteamagic_17-0.2.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_byteamagic_17` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_17-0.2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_byteamagic_17-0.2.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.7 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.7 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-byteamagic` | `0.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 KiB | [postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-17-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_byteamagic_16` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_byteamagic_16-0.2.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_byteamagic_16` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_byteamagic_16-0.2.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_byteamagic_16` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_byteamagic_16-0.2.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_byteamagic_16` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_byteamagic_16-0.2.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_byteamagic_16` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_byteamagic_16-0.2.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_byteamagic_16` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_16-0.2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_byteamagic_16-0.2.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.7 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.7 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.6 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-byteamagic` | `0.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.9 KiB | [postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-16-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_byteamagic_15` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_byteamagic_15-0.2.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_byteamagic_15` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_byteamagic_15-0.2.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_byteamagic_15` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_byteamagic_15-0.2.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_byteamagic_15` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_byteamagic_15-0.2.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_byteamagic_15` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_byteamagic_15-0.2.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_byteamagic_15` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_15-0.2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_byteamagic_15-0.2.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.7 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.7 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.6 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-byteamagic` | `0.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 KiB | [postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-15-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_byteamagic_14` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_byteamagic_14-0.2.4-1PIGSTY.el8.x86_64.rpm) |
| `pg_byteamagic_14` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_byteamagic_14-0.2.4-1PIGSTY.el8.aarch64.rpm) |
| `pg_byteamagic_14` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_byteamagic_14-0.2.4-1PIGSTY.el9.x86_64.rpm) |
| `pg_byteamagic_14` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_byteamagic_14-0.2.4-1PIGSTY.el9.aarch64.rpm) |
| `pg_byteamagic_14` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.8 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_byteamagic_14-0.2.4-1PIGSTY.el10.x86_64.rpm) |
| `pg_byteamagic_14` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.9 KiB | [pg_byteamagic_14-0.2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_byteamagic_14-0.2.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.7 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.6 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.6 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-byteamagic` | `0.2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 KiB | [postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-byteamagic/postgresql-14-pg-byteamagic_0.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/nmandery/pg_byteamagic" title="Repository" icon="github" subtitle="github.com/nmandery/pg_byteamagic" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_byteamagic-0.2.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_byteamagic;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_byteamagic;		# install via package name, for the active PG version
pig install byteamagic;		# install by extension name, for the current active PG version

pig install byteamagic -v 18;   # install for PG 18
pig install byteamagic -v 17;   # install for PG 17
pig install byteamagic -v 16;   # install for PG 16
pig install byteamagic -v 15;   # install for PG 15
pig install byteamagic -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION byteamagic;
```

## Usage

Sources: [official repo](https://github.com/nmandery/pg_byteamagic), [official doc](https://raw.githubusercontent.com/nmandery/pg_byteamagic/master/doc/byteamagic.md)

`byteamagic` runs `libmagic` on `bytea` values, so PostgreSQL can identify the MIME type or human-readable file type of blobs stored in the database. The package name is `pg_byteamagic`, but the extension name is `byteamagic`.

```sql
CREATE EXTENSION byteamagic;

SELECT byteamagic_mime(data) FROM files;
SELECT byteamagic_text(data) FROM files;
```

### Functions

- `byteamagic_mime(bytea) returns text`: MIME type, equivalent to `file --mime-type`.
- `byteamagic_text(bytea) returns text`: descriptive file type text, equivalent to `file`.

### Common Use

```sql
SELECT
  id,
  byteamagic_mime(blob) AS mime,
  byteamagic_text(blob) AS kind
FROM uploads;
```

### Caveats

- It only inspects `bytea` content; there are no operators, custom types, or extra SQL management objects.
- Build/install requires PostgreSQL development headers plus the system `libmagic` development package.
- The upstream documentation is minimal; current user-facing behavior is unchanged from the long-standing doc page.
