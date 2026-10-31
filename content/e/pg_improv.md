---
title: "pg_improv"
linkTitle: "pg_improv"
description: "Incremental multidimensional models and formula functions over PostgreSQL tables."
weight: 2580
categories: ["OLAP"]
languages: ["Rust"]
licenses: ["Apache-2.0 OR MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_improv**](https://codeberg.org/gregburd/improv) : Incremental multidimensional models and formula functions over PostgreSQL tables.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2580** | {{< badge content="pg_improv" link="https://codeberg.org/gregburd/improv" >}} | {{< ext "pg_improv" >}} | `0.12.2` | {{< category "OLAP" >}} | {{< license "Apache-2.0 OR MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `improv` |

> [!Note] Superuser install; commit mode needs preload; logical mode also needs shared preload/restart, wal_level=logical and a worker database. Logical slots retain WAL; numeric results use float8. PG15-18; build pgrx 0.19.3, upstream 0.17.0. 0.12.2 fixes repeated model-catalog reads by idle logical-sync workers.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_improv` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.2` | {{< bg "18" "pg_improv_18" "green" >}} {{< bg "17" "pg_improv_17" "green" >}} {{< bg "16" "pg_improv_16" "green" >}} {{< bg "15" "pg_improv_15" "green" >}} {{< bg "14" "pg_improv_14" "red" >}} | `pg_improv_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.1` | {{< bg "18" "postgresql-18-pg-improv" "green" >}} {{< bg "17" "postgresql-17-pg-improv" "green" >}} {{< bg "16" "postgresql-16-pg-improv" "green" >}} {{< bg "15" "postgresql-15-pg-improv" "green" >}} {{< bg "14" "postgresql-14-pg-improv" "red" >}} | `postgresql-$v-pg-improv` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.2" "pg_improv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_improv_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pg-improv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pg-improv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-improv : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_improv_18` | `0.12.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_improv_18-0.12.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_improv_18-0.12.2-1PGSTY.el8.x86_64.rpm) |
| `pg_improv_18` | `0.12.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.5 MiB | [pg_improv_18-0.12.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_improv_18-0.12.2-1PGSTY.el8.aarch64.rpm) |
| `pg_improv_18` | `0.12.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.2 MiB | [pg_improv_18-0.12.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_improv_18-0.12.2-1PGSTY.el9.x86_64.rpm) |
| `pg_improv_18` | `0.12.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.8 MiB | [pg_improv_18-0.12.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_improv_18-0.12.2-1PGSTY.el9.aarch64.rpm) |
| `pg_improv_18` | `0.12.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.2 MiB | [pg_improv_18-0.12.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_improv_18-0.12.2-1PGSTY.el10.x86_64.rpm) |
| `pg_improv_18` | `0.12.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [pg_improv_18-0.12.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_improv_18-0.12.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-improv` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.4 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.7 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.4 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.7 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.8 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.8 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.8 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-improv` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.3 MiB | [postgresql-18-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-18-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_improv_17` | `0.12.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_improv_17-0.12.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_improv_17-0.12.2-1PGSTY.el8.x86_64.rpm) |
| `pg_improv_17` | `0.12.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.5 MiB | [pg_improv_17-0.12.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_improv_17-0.12.2-1PGSTY.el8.aarch64.rpm) |
| `pg_improv_17` | `0.12.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.2 MiB | [pg_improv_17-0.12.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_improv_17-0.12.2-1PGSTY.el9.x86_64.rpm) |
| `pg_improv_17` | `0.12.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.8 MiB | [pg_improv_17-0.12.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_improv_17-0.12.2-1PGSTY.el9.aarch64.rpm) |
| `pg_improv_17` | `0.12.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.2 MiB | [pg_improv_17-0.12.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_improv_17-0.12.2-1PGSTY.el10.x86_64.rpm) |
| `pg_improv_17` | `0.12.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [pg_improv_17-0.12.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_improv_17-0.12.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-improv` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.4 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.7 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.4 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.7 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.8 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.8 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.8 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-improv` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.3 MiB | [postgresql-17-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-17-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_improv_16` | `0.12.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_improv_16-0.12.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_improv_16-0.12.2-1PGSTY.el8.x86_64.rpm) |
| `pg_improv_16` | `0.12.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.5 MiB | [pg_improv_16-0.12.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_improv_16-0.12.2-1PGSTY.el8.aarch64.rpm) |
| `pg_improv_16` | `0.12.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.2 MiB | [pg_improv_16-0.12.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_improv_16-0.12.2-1PGSTY.el9.x86_64.rpm) |
| `pg_improv_16` | `0.12.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.8 MiB | [pg_improv_16-0.12.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_improv_16-0.12.2-1PGSTY.el9.aarch64.rpm) |
| `pg_improv_16` | `0.12.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.2 MiB | [pg_improv_16-0.12.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_improv_16-0.12.2-1PGSTY.el10.x86_64.rpm) |
| `pg_improv_16` | `0.12.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.7 MiB | [pg_improv_16-0.12.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_improv_16-0.12.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-improv` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.4 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.7 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.4 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.7 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.8 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.3 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.8 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.3 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.8 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-improv` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.3 MiB | [postgresql-16-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-16-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_improv_15` | `0.12.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 5.3 MiB | [pg_improv_15-0.12.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_improv_15-0.12.2-1PGSTY.el8.x86_64.rpm) |
| `pg_improv_15` | `0.12.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 4.5 MiB | [pg_improv_15-0.12.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_improv_15-0.12.2-1PGSTY.el8.aarch64.rpm) |
| `pg_improv_15` | `0.12.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.2 MiB | [pg_improv_15-0.12.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_improv_15-0.12.2-1PGSTY.el9.x86_64.rpm) |
| `pg_improv_15` | `0.12.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 4.8 MiB | [pg_improv_15-0.12.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_improv_15-0.12.2-1PGSTY.el9.aarch64.rpm) |
| `pg_improv_15` | `0.12.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.2 MiB | [pg_improv_15-0.12.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_improv_15-0.12.2-1PGSTY.el10.x86_64.rpm) |
| `pg_improv_15` | `0.12.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 4.8 MiB | [pg_improv_15-0.12.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_improv_15-0.12.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-improv` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.4 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.7 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.4 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.7 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.4 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.4 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.8 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-improv` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.3 MiB | [postgresql-15-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-improv/postgresql-15-pg-improv_0.12.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/improv" title="Repository" icon="link" subtitle="codeberg.org/gregburd/improv" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="improv-0.12.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_improv;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_improv;		# install via package name, for the active PG version

pig install pg_improv -v 18;   # install for PG 18
pig install pg_improv -v 17;   # install for PG 17
pig install pg_improv -v 16;   # install for PG 16
pig install pg_improv -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_improv;
```

## Usage

Sources:

- [Release changelog](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.2/CHANGELOG.md)

- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.2/crates/pg_improv/README.md)
- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.2/crates/pg_improv/pg_improv.control)
- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.2/crates/pg_improv/Cargo.toml)

`pg_improv` 0.12.2 binds PostgreSQL tables to multidimensional models, evaluates formulas and materializes results. It supports PostgreSQL 15–18 and requires superuser installation. This example uses manual synchronization, which works without preloading the library.

### Core Workflow

```sql
CREATE EXTENSION pg_improv;
CREATE TABLE sales (quarter text, price float8, qty float8);
INSERT INTO sales VALUES ('Q1', 10, 100), ('Q2', 12, 200);
SELECT improv.create_model('plan');
SELECT improv.set_sync_mode('plan', 'manual');
SELECT improv.bind_category('plan', 'Time', 'sales', 'quarter');
SELECT improv.bind_measure('plan', 'Price', 'sales', 'price');
SELECT improv.bind_measure('plan', 'Quantity', 'sales', 'qty');
SELECT improv.define('plan', 'Revenue = Price * Quantity');
SELECT improv.recalc('plan');
SELECT * FROM improv_plan.revenue;
SELECT * FROM improv.status('plan');
```

Bind categories before their measures. Duplicate coordinates require an explicit aggregate; null values do not create model cells. Result tables live in a model-specific schema. `improv.define` accepts `replace => true` to replace a formula, and `improv.recalc` returns the number of changed result rows.

### Synchronization Modes

`improv.set_sync_mode` chooses manual, trigger, commit or logical synchronization. Trigger mode works without preloading; commit mode needs session or server preloading. The default logical mode needs `shared_preload_libraries`, a server restart, `wal_level=logical` and the appropriate `improv.database` worker setting. Logical results arrive asynchronously; use explicit recalculation when a transaction must immediately see current results.

Monitor `improv.status` and cap retained WAL with `max_slot_wal_keep_size`. Logical replication slots can retain WAL while a worker stalls. Binding tables can create publications and change replica identity; account for those effects before enabling logical mode.

### Formula and Maintenance Boundaries

The `improv` language supports scalar formulas as well as model expressions, but not arbitrary SQL calls or external I/O. Formula errors normally become null; `improv.raise_errors` changes that policy. Numeric values and large integers pass through floating-point evaluation, so precision beyond 2^53 needs care. `improv.max_cells` limits loaded input cells.

Back up model definitions and result state with database dumps and rehearse restoration; logical mode needs a fresh slot and recomputation after restore. Library upgrades require restarting preloaded processes and updating the extension SQL. Model computation is distinct from an application transaction's source-of-truth tables.

### Ownership and Database-Wide Logical Sync since 0.11.0

Each model now reads and writes as its owner; row-level security protects its catalog, items and cells. Only a superuser or a role allowed to act as the owner may change or recalculate that model. Grant improv_user to a model creator and grant the source-table privileges it needs; the administrator role does not automatically become every model owner. A shared result table still needs ordinary grants. what_if can reveal owner-visible data, so grant access deliberately. Logical models in one database now share one slot and publication. Upgrade with the matching library and SQL scripts, restore roles before model dumps, and monitor retained WAL.

### 0.12.2 Release Note

Release 0.12.2 fixes idle logical-sync workers so they no longer reread the model catalog on every poll; there are no other code changes in this release. The Pigsty build uses pgrx 0.19.3 for the current package, while upstream source metadata still declares pgrx 0.17.0.

### 0.12.1 Release Note

The 0.12.0 SQL migration adds `improv.measure.format` for persisted display formatting; update the extension before saving formats. Release 0.12.1 changes the version from 0.12.0 without additional code changes and restores RHEL 8 (glibc 2.28) compatible artifacts. Upstream still uses pgrx 0.17.0. Logical synchronization still requires `shared_preload_libraries`, `wal_level=logical` and a restart; trigger and manual modes do not.
