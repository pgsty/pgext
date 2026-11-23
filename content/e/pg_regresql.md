---
title: "pg_regresql"
linkTitle: "pg_regresql"
description: "Trust pg_class statistics for planning instead of physical relation size"
weight: 3230
categories: ["LANG"]
languages: ["C"]
licenses: ["BSD-2-Clause"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_regresql**](https://github.com/boringsql/regresql) : Trust pg_class statistics for planning instead of physical relation size


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3230** | {{< badge content="pg_regresql" link="https://github.com/boringsql/regresql" >}} | {{< ext "pg_regresql" >}} | `2.2.1` | {{< category "LANG" >}} | {{< license "BSD-2-Clause" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hypopg" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_hint_plan" >}} {{< ext "online_advisor" >}} {{< ext "index_advisor" >}} {{< ext "pg_stat_plans" >}} {{< ext "plan_filter" >}} {{< ext "pg_track_optimizer" >}} {{< ext "pg_strom" >}} {{< ext "pg_orca" >}} |

> [!Note] Load via LOAD or session_preload_libraries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_regresql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "19" "pg_regresql_19" "red" >}} {{< bg "18" "pg_regresql_18" "green" >}} {{< bg "17" "pg_regresql_17" "green" >}} {{< bg "16" "pg_regresql_16" "green" >}} {{< bg "15" "pg_regresql_15" "green" >}} | `pg_regresql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "19" "postgresql-19-pg-regresql" "red" >}} {{< bg "18" "postgresql-18-pg-regresql" "green" >}} {{< bg "17" "postgresql-17-pg-regresql" "green" >}} {{< bg "16" "postgresql-16-pg-regresql" "green" >}} {{< bg "15" "postgresql-15-pg-regresql" "green" >}} | `postgresql-$v-pg-regresql` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_regresql_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_regresql_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_regresql_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_regresql_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "pg_regresql_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "pg_regresql_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "pg_regresql_15 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-regresql : MISS 0" "red" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 90.5 KiB | [pg_regresql_18-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_18-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.5 KiB | [pg_regresql_18-2.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_18-2.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 88.3 KiB | [pg_regresql_18-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_18-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pg_regresql_18-2.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_18-2.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.0 KiB | [pg_regresql_18-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_18-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [pg_regresql_18-2.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_18-2.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.3 KiB | [pg_regresql_18-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_18-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [pg_regresql_18-2.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_18-2.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.3 KiB | [pg_regresql_18-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_18-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.5 KiB | [pg_regresql_18-2.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_18-2.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 89.3 KiB | [pg_regresql_18-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_18-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `pg_regresql_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [pg_regresql_18-2.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_18-2.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-regresql` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 75.4 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 73.2 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 75.6 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 73.3 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.4 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 79.8 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 78.5 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 76.7 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 77.8 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 76.2 KiB | [postgresql-18-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 106.6 KiB | [pg_regresql_17-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_17-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 KiB | [pg_regresql_17-2.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_17-2.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 103.8 KiB | [pg_regresql_17-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_17-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.5 KiB | [pg_regresql_17-2.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_17-2.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 107.7 KiB | [pg_regresql_17-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_17-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [pg_regresql_17-2.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_17-2.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.6 KiB | [pg_regresql_17-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_17-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [pg_regresql_17-2.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_17-2.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 108.0 KiB | [pg_regresql_17-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_17-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.5 KiB | [pg_regresql_17-2.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_17-2.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.6 KiB | [pg_regresql_17-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_17-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `pg_regresql_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [pg_regresql_17-2.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_17-2.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-regresql` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.2 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 88.0 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 91.5 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.5 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 116.5 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 114.4 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 95.0 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 93.1 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 94.7 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 92.5 KiB | [postgresql-17-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 106.5 KiB | [pg_regresql_16-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_16-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 KiB | [pg_regresql_16-2.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_16-2.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 103.4 KiB | [pg_regresql_16-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_16-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.5 KiB | [pg_regresql_16-2.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_16-2.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 107.5 KiB | [pg_regresql_16-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_16-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [pg_regresql_16-2.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_16-2.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 105.2 KiB | [pg_regresql_16-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_16-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [pg_regresql_16-2.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_16-2.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 107.8 KiB | [pg_regresql_16-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_16-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.5 KiB | [pg_regresql_16-2.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_16-2.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 105.2 KiB | [pg_regresql_16-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_16-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `pg_regresql_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [pg_regresql_16-2.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_16-2.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-regresql` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 91.1 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 87.8 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 91.3 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 88.1 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 114.8 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 112.4 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 95.0 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 92.7 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 94.6 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 92.3 KiB | [postgresql-16-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.5 KiB | [pg_regresql_15-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_15-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 KiB | [pg_regresql_15-2.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_15-2.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_regresql_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.7 KiB | [pg_regresql_15-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_15-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.5 KiB | [pg_regresql_15-2.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_15-2.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_regresql_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.7 KiB | [pg_regresql_15-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_15-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.4 KiB | [pg_regresql_15-2.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_15-2.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_regresql_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.7 KiB | [pg_regresql_15-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_15-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.3 KiB | [pg_regresql_15-2.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_15-2.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_regresql_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.7 KiB | [pg_regresql_15-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_15-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.5 KiB | [pg_regresql_15-2.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_15-2.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_regresql_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.8 KiB | [pg_regresql_15-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_15-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `pg_regresql_15` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [pg_regresql_15-2.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_15-2.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-regresql` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.5 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.6 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.5 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.6 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.0 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.1 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.9 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.0 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.2 KiB | [postgresql-15-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/boringsql/regresql" title="Repository" icon="github" subtitle="github.com/boringsql/regresql" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_regresql-2.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_regresql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_regresql;		# install via package name, for the active PG version

pig install pg_regresql -v 19;   # install for PG 19
pig install pg_regresql -v 18;   # install for PG 18
pig install pg_regresql -v 17;   # install for PG 17
pig install pg_regresql -v 16;   # install for PG 16
pig install pg_regresql -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_regresql';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_regresql;
```

## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/README.md)
- [pg_regresql.control](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql.control)
- [Installation SQL](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql--2.2.1.sql)
- [Upgrade SQL from 2.2](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql--2.2--2.2.1.sql)

`pg_regresql` makes query planning use catalog statistics instead of rescaling them from physical file size. This supports plan regression tests in databases with representative statistics but little data. Release 2.2.1 installs its SQL objects in the fixed `pg_catalog` schema.

### Activate the Planner Hook

The hook supports PostgreSQL 15 or later. Load it as a superuser in the test session:

```sql
LOAD 'pg_regresql';
CREATE TABLE test_orders (created_at date);
EXPLAIN SELECT * FROM test_orders WHERE created_at > DATE '2024-06-01';
```

For new sessions on a test instance, add the library to the existing session preload list:

```ini
session_preload_libraries = 'pg_regresql'
```

The hook replaces table and index page/tuple estimates with `pg_class` values and computes visibility from catalog statistics. Inject useful statistics before comparing plans; an empty table alone is not a production workload.

### Restore Statistics

On PostgreSQL 16 and 17, create the extension before loading statistics emitted by a recent pg_dump:

```sql
CREATE EXTENSION pg_regresql;
SELECT pg_restore_relation_stats(
    'schemaname', 'public', 'relname', 'test_orders',
    'relpages', 123513::integer, 'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);
```

Version 2.2 adds `pg_restore_extended_stats()` and `pg_clear_extended_stats()` for those servers, and supports the JSON input emitted by pg_dump 19 for `pg_ndistinct` and `pg_dependencies`. On PostgreSQL 16 remove the dump's unsupported transaction_timeout setting. PostgreSQL 18 already supplies relation/attribute restoration in core; extended restoration is in core from 19. The extension does not install these backports on 18 or later.

### Upgrade and Cleanup

From SQL version 2.0, upstream requires dropping and recreating the extension. From 2.1, use `ALTER EXTENSION pg_regresql UPDATE`. Review dependencies before removing it. The extension is not relocatable and its install requires a superuser.

The replacement input functions named `regresql_pg_*_in` are deliberately detached from the extension on 16/17. `DROP EXTENSION` leaves them behind so the installed casts keep working; dropping the extension is therefore not a complete return to the original catalog state. Use this tool in controlled planning tests, and do not assume compatibility with other planner hooks has been tested.

### PostgreSQL 19 Planner Hook

Release 2.2.1 uses `build_simple_rel_hook` on PostgreSQL 19 and `get_relation_info_hook` on earlier supported majors. After installing the matching files, run `ALTER EXTENSION pg_regresql UPDATE TO '2.2.1';` on installations using 2.2. This update advances the recorded extension version and preserves the existing SQL definitions.
