---
title: "pg_try_cast"
linkTitle: "pg_try_cast"
description: "Text conversion with NULL or a typed fallback on conversion errors."
weight: 4855
categories: ["FUNC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_try_cast**](https://github.com/Mazuco/pg_try_cast) : Text conversion with NULL or a typed fallback on conversion errors.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4855** | {{< badge content="pg_try_cast" link="https://github.com/Mazuco/pg_try_cast" >}} | {{< ext "pg_try_cast" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |

> [!Note] Release 1.0.0 installs control/SQL 1.0; META provides.version says 1.0.0. PG16+ required, CI covers 16-18; input types without soft-error support may still raise errors.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_try_cast` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_try_cast_18" "green" >}} {{< bg "17" "pg_try_cast_17" "green" >}} {{< bg "16" "pg_try_cast_16" "green" >}} {{< bg "15" "pg_try_cast_15" "red" >}} {{< bg "14" "pg_try_cast_14" "red" >}} | `pg_try_cast_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-try-cast" "green" >}} {{< bg "17" "postgresql-17-pg-try-cast" "green" >}} {{< bg "16" "postgresql-16-pg-try-cast" "green" >}} {{< bg "15" "postgresql-15-pg-try-cast" "red" >}} {{< bg "14" "postgresql-14-pg-try-cast" "red" >}} | `postgresql-$v-pg-try-cast` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_try_cast_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_try_cast_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_try_cast_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-try-cast : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-try-cast : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-try-cast : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_try_cast_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.4 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_try_cast_18-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_try_cast_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.4 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_try_cast_18-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_try_cast_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.7 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_try_cast_18-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_try_cast_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_try_cast_18-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_try_cast_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_try_cast_18-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_try_cast_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_try_cast_18-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_try_cast_18-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.7 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.7 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.7 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.7 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.8 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.7 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-try-cast` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.8 KiB | [postgresql-18-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-18-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_try_cast_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.4 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_try_cast_17-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_try_cast_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.4 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_try_cast_17-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_try_cast_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.7 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_try_cast_17-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_try_cast_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.4 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_try_cast_17-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_try_cast_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_try_cast_17-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_try_cast_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_try_cast_17-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_try_cast_17-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.6 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.7 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.6 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.7 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.7 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.9 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.8 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.7 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-try-cast` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.8 KiB | [postgresql-17-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-17-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_try_cast_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.4 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_try_cast_16-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `pg_try_cast_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.4 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_try_cast_16-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `pg_try_cast_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.7 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_try_cast_16-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `pg_try_cast_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.4 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_try_cast_16-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `pg_try_cast_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_try_cast_16-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `pg_try_cast_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_try_cast_16-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_try_cast_16-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.6 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.7 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.6 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.8 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.7 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.9 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.8 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.7 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.9 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-try-cast` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.8 KiB | [postgresql-16-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-try-cast/postgresql-16-pg-try-cast_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Mazuco/pg_try_cast" title="Repository" icon="github" subtitle="github.com/Mazuco/pg_try_cast" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_try_cast-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_try_cast;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_try_cast;		# install via package name, for the active PG version

pig install pg_try_cast -v 18;   # install for PG 18
pig install pg_try_cast -v 17;   # install for PG 17
pig install pg_try_cast -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_try_cast;
```

## Usage

Sources:

- [README.md](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/README.md)
- [pg_try_cast--1.0.sql](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast--1.0.sql)
- [src/pg_try_cast.c](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/src/pg_try_cast.c)
- [.github/workflows/ci.yml](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/.github/workflows/ci.yml)
- [pg_try_cast.control](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast.control)

`pg_try_cast` provides `try_cast` and `try_convert` for converting text without aborting a query on supported input errors. Release 1.0.0 installs SQL version 1.0. PostgreSQL 16 or later is required; upstream CI covers 16–18. It is trusted and relocatable, with no preload requirement.

### Core Workflow

```sql
CREATE EXTENSION pg_try_cast;
SELECT try_cast('42', NULL::integer);
SELECT try_cast('invalid', 0);
SELECT try_convert(NULL::date, '2026-02-30');
```

### Operational Boundaries

`try_cast(value text, fallback anyelement)` takes its result type and error fallback from the second argument; `try_convert(fallback anyelement, value text)` reverses the order. NULL input returns NULL. Give the fallback an explicit type. Type modifiers and domain constraints are checked. This covers text input, not arbitrary native casts. Input functions without soft-error support can still raise an error, including malformed type names passed to regtype. A role with CREATE privilege on the database can install this trusted extension after its files are installed.
