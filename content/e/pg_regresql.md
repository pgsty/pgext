---
title: "pg_regresql"
linkTitle: "pg_regresql"
description: "Trust pg_class statistics for planning instead of physical relation size"
weight: 3230
categories: ["LANG"]
width: full
---

[**pg_regresql**](https://github.com/boringsql/regresql) : Trust pg_class statistics for planning instead of physical relation size


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3230** | {{< badge content="pg_regresql" link="https://github.com/boringsql/regresql" >}} | {{< ext "pg_regresql" >}} | `2.0.0` | {{< category "LANG" >}} | {{< license "BSD-2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} {{< ext "plan_filter" >}} {{< ext "auto_explain" >}} |

> [!Note] Activate it with LOAD pg_regresql or session_preload_libraries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_regresql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "pg_regresql_18" "green" >}} {{< bg "17" "pg_regresql_17" "green" >}} {{< bg "16" "pg_regresql_16" "green" >}} {{< bg "15" "pg_regresql_15" "green" >}} {{< bg "14" "pg_regresql_14" "green" >}} | `pg_regresql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "postgresql-18-pg-regresql" "green" >}} {{< bg "17" "postgresql-17-pg-regresql" "green" >}} {{< bg "16" "postgresql-16-pg-regresql" "green" >}} {{< bg "15" "postgresql-15-pg-regresql" "green" >}} {{< bg "14" "postgresql-14-pg-regresql" "green" >}} | `postgresql-$v-pg-regresql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_regresql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-pg-regresql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-pg-regresql : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.7 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_18-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_regresql_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.9 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_18-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_regresql_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.4 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_18-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_regresql_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.4 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_18-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_regresql_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_18-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_regresql_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_regresql_18-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_18-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-regresql` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.1 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.1 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.6 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.3 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.4 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-regresql` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.3 KiB | [postgresql-18-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-18-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.7 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_regresql_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.9 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_regresql_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.4 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_regresql_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.4 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_regresql_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_17-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_regresql_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_regresql_17-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_17-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-regresql` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.1 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.1 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.4 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.3 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-regresql` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.3 KiB | [postgresql-17-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-17-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.7 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_regresql_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.9 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_regresql_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.4 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_regresql_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.4 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_regresql_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_16-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_regresql_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_regresql_16-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_16-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-regresql` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.1 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.1 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.4 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.3 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-regresql` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.3 KiB | [postgresql-16-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-16-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.7 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_15-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_regresql_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.9 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_15-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_regresql_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.4 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_15-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_regresql_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.4 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_15-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_regresql_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_15-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_regresql_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_regresql_15-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_15-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-regresql` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.1 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.1 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.3 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.0 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.3 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-regresql` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.3 KiB | [postgresql-15-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-15-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_regresql_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.7 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_regresql_14-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_regresql_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.9 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_regresql_14-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_regresql_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.4 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_regresql_14-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_regresql_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.4 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_regresql_14-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_regresql_14` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_regresql_14-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_regresql_14` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_regresql_14-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_regresql_14-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-regresql` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.1 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.1 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.1 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.3 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.0 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.3 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-regresql` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.3 KiB | [postgresql-14-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-regresql/postgresql-14-pg-regresql_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/boringsql/regresql" title="Repository" icon="github" subtitle="github.com/boringsql/regresql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_regresql-2.0.0.zip" >}}
{{< /cards >}}


```bash
pig build pkg pg_regresql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_regresql;		# install via package name, for the active PG version

pig install pg_regresql -v 18;   # install for PG 18
pig install pg_regresql -v 17;   # install for PG 17
pig install pg_regresql -v 16;   # install for PG 16
pig install pg_regresql -v 15;   # install for PG 15
pig install pg_regresql -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_regresql';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_regresql;
```


## Usage

> Syntax:
>
> ```bash
> regresql init postgres://localhost/mydb
> regresql add src/sql/
> regresql update
> regresql test
> ```
>
> Sources: [README](https://github.com/boringsql/regresql), [Product page](https://boringsql.com/products/regresql/)

`RegreSQL` is documented upstream as a language-agnostic SQL regression testing tool for PostgreSQL, not as a `CREATE EXTENSION`-style in-database module. It discovers `.sql` files, runs them against PostgreSQL, snapshots expected output, and tracks query plan changes.

## Quick Start

The README's basic workflow is:

```bash
regresql init postgres://localhost/mydb
regresql discover
regresql add src/sql/
regresql update
regresql test
```

This initializes a test suite, discovers query files, creates plan definitions, captures expected output, and runs regression checks.

## What It Tracks

The upstream docs emphasize:

- expected query output snapshots
- `EXPLAIN` plan baselines
- sequential scan warnings
- migration-related query regressions
- CI-oriented output formats such as `junit`, `json`, `pgtap`, and `github-actions`

## Query Files and Plans

RegreSQL works with normal SQL files and supports multiple queries per file using `-- name:` annotations:

```sql
-- name: get-user-by-id
SELECT * FROM users WHERE id = :id;
```

Plan files provide test parameters:

```yaml
"1":
  id: 42
"2":
  id: 100
```

## Snapshots and Migrations

The tool can build and restore database snapshots and compare query behavior across migrations:

```bash
regresql snapshot build
regresql snapshot restore
regresql migrate --script db/migrations/001_add_column.sql
```

## Installation

The README documents installation via Homebrew or Go:

```bash
brew tap boringsql/boringsql
brew install regresql
```

or

```bash
go install github.com/boringsql/regresql@latest
```

PostgreSQL client tools such as `pg_dump`, `pg_restore`, and `psql` are required for snapshot commands.
