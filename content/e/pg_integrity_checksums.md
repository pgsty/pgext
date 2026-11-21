---
title: "pg_integrity_checksums"
linkTitle: "pg_integrity_checksums"
description: "Physical and logical checksums for PostgreSQL pages, rows, tables, indexes and databases."
weight: 5115
categories: ["ADMIN"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_integrity_checksums**](https://github.com/VictorNagibator/pg_integrity_checksums) : Physical and logical checksums for PostgreSQL pages, rows, tables, indexes and databases.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5115** | {{< badge content="pg_integrity_checksums" link="https://github.com/VictorNagibator/pg_integrity_checksums" >}} | {{< ext "pg_integrity_checksums" >}} | `2.9.1` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "pg_checksums" >}} |

> [!Note] 2.2.0 adds FNV-1a 64-bit overloads; superuser; recompute affected baselines.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.9.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_integrity_checksums` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.9.1` | {{< bg "19" "pg_integrity_checksums_19" "green" >}} {{< bg "18" "pg_integrity_checksums_18" "green" >}} {{< bg "17" "pg_integrity_checksums_17" "green" >}} {{< bg "16" "pg_integrity_checksums_16" "green" >}} {{< bg "15" "pg_integrity_checksums_15" "green" >}} | `pg_integrity_checksums_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.9.1` | {{< bg "19" "postgresql-19-pg-integrity-checksums" "green" >}} {{< bg "18" "postgresql-18-pg-integrity-checksums" "green" >}} {{< bg "17" "postgresql-17-pg-integrity-checksums" "green" >}} {{< bg "16" "postgresql-16-pg-integrity-checksums" "green" >}} {{< bg "15" "postgresql-15-pg-integrity-checksums" "green" >}} | `postgresql-$v-pg-integrity-checksums` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_integrity_checksums_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_integrity_checksums_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_integrity_checksums_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_integrity_checksums_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "pg_integrity_checksums_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "pg_integrity_checksums_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_integrity_checksums_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-integrity-checksums : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-integrity-checksums : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-integrity-checksums : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_integrity_checksums_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 109.3 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_integrity_checksums_18-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_integrity_checksums_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.6 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_integrity_checksums_18-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_integrity_checksums_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.8 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_integrity_checksums_18-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_integrity_checksums_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 111.7 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_integrity_checksums_18-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_integrity_checksums_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.5 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_integrity_checksums_18-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_integrity_checksums_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.2 KiB | [pg_integrity_checksums_18-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_integrity_checksums_18-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.8 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 97.8 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.1 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 97.9 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 109.0 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 106.6 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 105.7 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 104.1 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 104.2 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-integrity-checksums` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 101.8 KiB | [postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-18-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_integrity_checksums_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 109.1 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_integrity_checksums_17-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_integrity_checksums_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.5 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_integrity_checksums_17-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_integrity_checksums_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.4 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_integrity_checksums_17-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_integrity_checksums_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 111.5 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_integrity_checksums_17-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_integrity_checksums_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 114.3 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_integrity_checksums_17-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_integrity_checksums_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.1 KiB | [pg_integrity_checksums_17-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_integrity_checksums_17-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.6 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 97.7 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.1 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 97.9 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.0 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.4 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 105.5 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 104.0 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 104.0 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-integrity-checksums` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 101.8 KiB | [postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-17-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_integrity_checksums_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 110.0 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_integrity_checksums_16-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_integrity_checksums_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 107.1 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_integrity_checksums_16-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_integrity_checksums_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 114.2 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_integrity_checksums_16-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_integrity_checksums_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 112.0 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_integrity_checksums_16-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_integrity_checksums_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.0 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_integrity_checksums_16-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_integrity_checksums_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.8 KiB | [pg_integrity_checksums_16-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_integrity_checksums_16-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 101.2 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 98.3 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.6 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 98.4 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.6 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 114.0 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 106.2 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 104.6 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 104.6 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-integrity-checksums` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.4 KiB | [postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-16-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_integrity_checksums_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 109.7 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_integrity_checksums_15-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_integrity_checksums_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.8 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_integrity_checksums_15-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_integrity_checksums_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 113.8 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_integrity_checksums_15-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_integrity_checksums_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 111.8 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_integrity_checksums_15-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_integrity_checksums_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 115.0 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_integrity_checksums_15-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_integrity_checksums_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 112.3 KiB | [pg_integrity_checksums_15-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_integrity_checksums_15-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 100.5 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 97.5 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 101.0 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 97.7 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.1 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.6 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 105.6 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 103.9 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 104.4 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-integrity-checksums` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 102.0 KiB | [postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-integrity-checksums/postgresql-15-pg-integrity-checksums_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/VictorNagibator/pg_integrity_checksums" title="Repository" icon="github" subtitle="github.com/VictorNagibator/pg_integrity_checksums" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_integrity_checksums-2.9.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_integrity_checksums;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_integrity_checksums;		# install via package name, for the active PG version

pig install pg_integrity_checksums -v 19;   # install for PG 19
pig install pg_integrity_checksums -v 18;   # install for PG 18
pig install pg_integrity_checksums -v 17;   # install for PG 17
pig install pg_integrity_checksums -v 16;   # install for PG 16
pig install pg_integrity_checksums -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_integrity_checksums;
```

## Usage

Sources:

- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/README.md)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums.control)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/Changes)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums--2.9.0--2.9.1.sql)

- [2.9.0 Changes](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/Changes)
- [2.8.0 to 2.9.0 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums--2.8.0--2.9.0.sql)
- [2.9.0 control](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums.control)
- [2.9.0 function reference](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/docs/functions.md)
- [Current upstream](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/Changes)
- [Current upstream](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.6.0--2.7.0.sql)

- [PGXN metadata](https://api.pgxn.org/dist/pg_integrity_checksums/2.7.0/META.json)
- [Official README](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/README.md)
- [Control file](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums.control)
- [Function reference](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/functions.md)
- [Operations guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/operations.md)
- [2.5 to 2.6 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.5.0--2.6.0.sql)
- [2.4 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.3.0--2.4.0.sql)
- [2.5 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.4.0--2.5.0.sql)
- [Migration guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/MIGRATION.md)
- [Upgrade and baseline guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/installation.md)

`pg_integrity_checksums` computes physical and logical checksums for PostgreSQL pages, cells, tuples, tables, indexes and databases. This C extension supports PostgreSQL 12–19 upstream; the catalog's Pigsty packages target PostgreSQL 14–18. Installation and database-wide checks require a superuser.

### Install and migrate

```sql
CREATE EXTENSION pg_integrity_checksums;
```

### Version 2.3.0 and page verification

Version 2.3.0 adds `pg_page_checksum_status(regclass, integer)`, which compares a page's calculated 16-bit checksum with its stored `pd_checksum`. The following checks remain available after the upgrade:

```sql
SHOW data_checksums;
SELECT pg_page_checksum_status('public.orders'::regclass, 0);
```

The relation must exist and contain the requested block. Results are `ok` for a match, `mismatch` for a difference, `new` for an uninitialized page, or `checksums_disabled` when stored page checksums are disabled. Stored-page verification requires `data_checksums = on`; an invalid relation or block raises an error.

`pg_page_checksum()` only calculates PostgreSQL's 16-bit page checksum; it does not compare the stored value. Its result of zero can mean an uninitialized page or a calculated zero, so zero alone does not establish validity. The status function reads relation storage directly and reports the page it reads, rather than a persistent guarantee about later writes.

The FNV-1a 32/64-bit overloads introduced in 2.2.0 still return raw `bytea` digests alongside the legacy 32-bit SQL API. They are separate from PostgreSQL's 16-bit page-checksum format. Changing the selected algorithm or overload requires a matching baseline; the 2.3.0 page-status addition does not itself require recomputing all existing baselines.

### Migration from the old extension name

Version 2.0.0 is the renamed continuation of the earlier `pg_checksums` extension. PostgreSQL has no `ALTER EXTENSION ... RENAME`; review dependent objects, drop the old extension, create the new one, and change the GUC prefix before using the new name:

```sql
DROP EXTENSION pg_checksums;
CREATE EXTENSION pg_integrity_checksums;
SET pg_integrity_checksums.workers = 4;
SET pg_integrity_checksums.strict = on;
```

The old and new names cannot coexist because they define the same SQL functions. The unrelated core `pg_checksums` utility is a separate program.

### Checksum levels

Use logical checksums for content comparison and physical checksums for storage integrity. Logical table checksums do not require a primary key, preserve duplicate rows, and exclude relation and tuple identifiers. Logical index checksums are derived from visible table rows and the index definition, so they are stable across `VACUUM`, `CLUSTER` and `REINDEX`.

```sql
SELECT pg_table_logical_checksum('public.orders'::regclass);
SELECT pg_index_logical_checksum('public.orders_pkey'::regclass);
SELECT pg_database_logical_checksum(false, true);
```

Physical checksums include storage details and can change after `VACUUM FULL`, `CLUSTER`, `REINDEX` or other physical rewrites. Use the page, tuple and table physical functions when those changes are relevant to the integrity baseline.

### Privileges and maintenance

Relation-level functions require table-level `SELECT` and reject row-level-security protected relations because direct storage scans cannot apply RLS; column-only grants are insufficient. Database-level functions are superuser-only. `pg_integrity_checksums.workers` controls parallel scans; `pg_integrity_checksums.strict` makes database scans fail when a relation cannot be checked instead of skipping it. The extension does not require shared preload. Preserve a known-good baseline for the same algorithm and compatible schema, and account for physical rewrites when comparing physical checksums. These checksums are not cryptographic authentication.

### Upgrade and baseline changes in 2.7.0

Install matching files, reconnect and apply the SQL upgrade:

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.7.0';
SHOW pg_integrity_checksums.strict;
```

Database-level logical checksums now include qualified relation names, so every database logical baseline changes even if row contents do not. Some physical index baselines also change after index metadata masking was corrected. Recompute and review these baselines; physical checksums are still sensitive to real storage changes, including row locks and rolled-back writes. A checksum difference alone does not establish data corruption.

Tuple and cell checksums now check MVCC visibility and return NULL for invisible rows. Index expressions execute as the table owner in a restricted context, and serializable scans take the required conflict locks. Page-status checks retry disk mismatches; a relation dropped during a database scan is skipped with NOTICE even in strict mode. Other unreadable relations still follow `pg_integrity_checksums.strict`, which defaults to on. These semantics differ from the earlier 2.6.0 masking and baseline rules.

### 2.9.0 Algorithms, Sequences and Baselines

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.0';
```

Version 2.9.0 adds `crc32c` and `xxh3` to the generic algorithm functions and extends parallel workers to `pg_table_bucket_checksums()`. New sequence logical checksums compare a sequence's definition with `last_value` and `is_called`; new relation-page status functions verify all pages and forks in one call.

When `include_system = false`, database logical and physical checksums now exclude `information_schema` as well as `pg_catalog`. Because `information_schema` content varies by PostgreSQL version, every non-system database checksum from an older release changes after this correction: recompute and replace those stored baselines. Table, index and sequence checksum values are not changed by this fix, and CLI baselines already excluded `information_schema`.


### 2.9.1 Scan Changes

On PostgreSQL 17+, table, index, bucket and database scans use read streams; PostgreSQL 18 can use asynchronous I/O. Parallel workers process larger ranges. Version 2.9.1 does not change checksum values. Install the matching library and apply `ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.1'`; the older 2.9.0 information_schema exclusion still requires replacing affected non-system database baselines.
