---
title: "fbsql"
linkTitle: "fbsql"
description: "Closure-preserving formula-based statistical modeling in SQL"
weight: 4695
categories: ["FUNC"]
width: full
---

[**fbsql**](https://github.com/dsc-chiba-u/FbSQL) : Closure-preserving formula-based statistical modeling in SQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4695** | {{< badge content="fbsql" link="https://github.com/dsc-chiba-u/FbSQL" >}} | {{< ext "fbsql" >}} | `0.1.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fbsql` |
|   **Requires**    | {{< ext "plr" >}} |
|   **See Also**    | {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "pg_math" >}} {{< ext "weighted_statistics" >}} |

> [!Note] Requires PL/R 8.4.0 or newer; PIGSTY packages target PostgreSQL 16 through 18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `fbsql` | `plr` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "fbsql_18" "green" >}} {{< bg "17" "fbsql_17" "green" >}} {{< bg "16" "fbsql_16" "green" >}} {{< bg "15" "fbsql_15" "red" >}} {{< bg "14" "fbsql_14" "red" >}} | `fbsql_$v` | `plr_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-fbsql" "green" >}} {{< bg "17" "postgresql-17-fbsql" "green" >}} {{< bg "16" "postgresql-16-fbsql" "green" >}} {{< bg "15" "postgresql-15-fbsql" "red" >}} {{< bg "14" "postgresql-14-fbsql" "red" >}} | `postgresql-$v-fbsql` | `postgresql-$v-plr` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "fbsql_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "fbsql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "fbsql_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-fbsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-fbsql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-fbsql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-fbsql : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fbsql_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [fbsql_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fbsql_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 KiB | [fbsql_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fbsql_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.6 KiB | [fbsql_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fbsql_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [fbsql_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fbsql_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [fbsql_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fbsql_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `fbsql_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.7 KiB | [fbsql_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fbsql_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-fbsql` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.8 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.8 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-fbsql` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.9 KiB | [postgresql-18-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-18-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fbsql_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [fbsql_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fbsql_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 KiB | [fbsql_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fbsql_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.6 KiB | [fbsql_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fbsql_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [fbsql_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fbsql_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [fbsql_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fbsql_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `fbsql_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.7 KiB | [fbsql_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fbsql_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-fbsql` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.8 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.8 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-fbsql` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.9 KiB | [postgresql-17-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-17-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `fbsql_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [fbsql_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/fbsql_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.4 KiB | [fbsql_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/fbsql_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `fbsql_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.6 KiB | [fbsql_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/fbsql_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.5 KiB | [fbsql_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/fbsql_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `fbsql_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [fbsql_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/fbsql_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `fbsql_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.7 KiB | [fbsql_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/fbsql_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-fbsql` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.8 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.8 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-fbsql` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.9 KiB | [postgresql-16-fbsql_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/fbsql/postgresql-16-fbsql_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dsc-chiba-u/FbSQL" title="Repository" icon="github" subtitle="github.com/dsc-chiba-u/FbSQL" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="fbsql-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg fbsql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install fbsql;		# install via package name, for the active PG version

pig install fbsql -v 18;   # install for PG 18
pig install fbsql -v 17;   # install for PG 17
pig install fbsql -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION fbsql CASCADE; -- requires plr
```

## Usage

Sources:

- [FbSQL 0.1.0 README](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/README.md)
- [FbSQL 0.1.0 changes](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/Changes)
- [Extension control file](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/fbsql.control)
- [PGXN release](https://pgxn.org/dist/fbsql/0.1.0/)

`fbsql` is a proof-of-concept statistical-modeling DSL that keeps fitting and prediction relational: SQL queries go in and rows come back, while models are described with R formula syntax. Release 0.1.0 implements generalized linear models through PL/R for fitting and pure PL/pgSQL for prediction.

### Prerequisites

FbSQL was developed and tested with PostgreSQL 16 and requires PL/R 8.4.0 or newer plus R. `plr` is an untrusted language, so a superuser must install the dependency and extension.

```sql
CREATE EXTENSION fbsql CASCADE;
SELECT fbsql.version();
```

Grant regular users only the function access and source-data privileges they require.

### Core Workflow

Fit a binomial churn model and retain the returned relation:

```sql
CREATE TEMP TABLE churn_model AS
SELECT *
FROM fbsql.fit_glm(
  relation => $$
    SELECT churn_flag, age, gender
    FROM customer
    WHERE created_at >= DATE '2025-01-01'
      AND created_at <  DATE '2026-01-01'
  $$,
  formula => 'churn_flag ~ age + gender',
  family => 'binomial'
);
```

Prediction accepts a query for new rows and a query returning the saved model. Because it returns `SETOF record`, supply the output columns at the call site:

```sql
SELECT customer_id, churn_flag_predicted
FROM fbsql.predict_glm(
  relation => $$SELECT customer_id, age, gender FROM customer_2026$$,
  model    => $$SELECT * FROM churn_model$$
) AS p(
  customer_id bigint,
  age integer,
  gender text,
  churn_flag_predicted double precision
);
```

### Important Objects

- `fbsql.fit_glm(relation, formula, family)` returns one row per model term, repeated fit statistics, and `metadata jsonb` containing the information needed for prediction.
- `fbsql.predict_glm(relation, model, on_new_levels)` appends `<response>_predicted` to the input rows. `on_new_levels` is `error` by default or `na` to produce a null prediction for unseen factor levels.
- `fbsql.version()` reports the extension version.

### Supported Surface and Caveats

Version 0.1.0 supports Gaussian models with the identity link and binomial models with the logit link, using numeric and factor predictors. Fitting applies complete-case analysis and reports used and dropped row counts; prediction returns `NULL` when a predictor is null. Prediction uses stored coefficients and metadata and does not invoke R at runtime.

Interactions, custom contrasts, offsets, weights, prediction intervals, additional families and links, and distributed fitting are not supported. The `relation` and `model` parameters contain SQL text: construct them from trusted SQL, not unsanitized user input, and review the executing role's privileges.
