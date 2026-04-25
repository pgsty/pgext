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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_regresql-2.0.0.tar.gz" >}}
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

Sources: [extension README](https://github.com/boringSQL/regresql/blob/master/pg_ext/README.md), [control file](https://github.com/boringSQL/regresql/blob/master/pg_ext/pg_regresql.control), [portable stats article](https://boringsql.com/posts/portable-stats/)

`pg_regresql` is the PostgreSQL planner-hook part of RegreSQL. It makes the planner trust `pg_class` statistics instead of rescaling estimates from the physical file size, which is useful for plan regression testing with injected production statistics.

### Activate the hook

```sql
LOAD 'pg_regresql';

EXPLAIN SELECT ...;
```

For a whole test instance, upstream recommends:

```conf
session_preload_libraries = 'pg_regresql'
```

### What it overrides

The upstream README says the hook runs after `estimate_rel_size()` and replaces planner estimates with catalog values:

- `rel->pages` from `pg_class.relpages`
- `rel->tuples` from `pg_class.reltuples`
- `rel->allvisfrac` from `pg_class.relallvisible / relpages`
- `IndexOptInfo->pages` from index `pg_class.relpages`
- `IndexOptInfo->tuples` from index `pg_class.reltuples`

### Typical workflow

```sql
SELECT pg_restore_relation_stats(
    'schemaname', 'public',
    'relname', 'test_orders',
    'relpages', 123513::integer,
    'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);

LOAD 'pg_regresql';

EXPLAIN SELECT * FROM test_orders WHERE created_at > '2024-06-01';
```

This is meant for reproducing production plans locally and for migration or upgrade plan-regression tests.

### Notes

- The control file currently declares `default_version = '2.0'`.
- The public repository tags visible upstream are still `v2.0.0-alpha*`, so the packaged `2.0.0` target is ahead of a clearly tagged final GitHub release.
- Upstream documents PostgreSQL 14+ compatibility for the extension.
