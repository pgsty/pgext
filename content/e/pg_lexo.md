---
title: "pg_lexo"
linkTitle: "pg_lexo"
description: "Lexicographic position type for inserting and reordering list items"
weight: 3530
categories: ["TYPE"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_lexo**](https://github.com/Blad3Mak3r/pg_lexo) : Lexicographic position type for inserting and reordering list items


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3530** | {{< badge content="pg_lexo" link="https://github.com/Blad3Mak3r/pg_lexo" >}} | {{< ext "pg_lexo" >}} | `0.6.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "base62" >}} {{< ext "pg_hashids" >}} {{< ext "typeid" >}} {{< ext "sequential_uuids" >}} |

> [!Note] Release/package 0.6.1 ships SQL/control version 0.6.0. Objects use the chosen installation schema; the old lexo schema was removed. superuser=false is not trusted=true. PGSTY patches lexo_next for native-type ordering and empty tables, and lexo_rebalance to preserve order independently of text collation.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lexo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "pg_lexo_18" "green" >}} {{< bg "17" "pg_lexo_17" "green" >}} {{< bg "16" "pg_lexo_16" "green" >}} {{< bg "15" "pg_lexo_15" "red" >}} {{< bg "14" "pg_lexo_14" "red" >}} | `pg_lexo_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "postgresql-18-pg-lexo" "green" >}} {{< bg "17" "postgresql-17-pg-lexo" "green" >}} {{< bg "16" "postgresql-16-pg-lexo" "green" >}} {{< bg "15" "postgresql-15-pg-lexo" "red" >}} {{< bg "14" "postgresql-14-pg-lexo" "red" >}} | `postgresql-$v-pg-lexo` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pg_lexo_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lexo_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lexo_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pg-lexo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pg-lexo : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lexo : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lexo : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lexo_18` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 378.5 KiB | [pg_lexo_18-0.6.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lexo_18-0.6.1-1PGSTY.el8.x86_64.rpm) |
| `pg_lexo_18` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 265.6 KiB | [pg_lexo_18-0.6.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lexo_18-0.6.1-1PGSTY.el8.aarch64.rpm) |
| `pg_lexo_18` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 395.5 KiB | [pg_lexo_18-0.6.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lexo_18-0.6.1-1PGSTY.el9.x86_64.rpm) |
| `pg_lexo_18` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 284.9 KiB | [pg_lexo_18-0.6.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lexo_18-0.6.1-1PGSTY.el9.aarch64.rpm) |
| `pg_lexo_18` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 395.5 KiB | [pg_lexo_18-0.6.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lexo_18-0.6.1-1PGSTY.el10.x86_64.rpm) |
| `pg_lexo_18` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 285.2 KiB | [pg_lexo_18-0.6.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lexo_18-0.6.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-lexo` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 323.5 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 216.2 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 323.3 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 216.1 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 364.1 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 251.5 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 357.9 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 250.3 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 356.6 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-lexo` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 249.6 KiB | [postgresql-18-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-18-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lexo_17` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 378.4 KiB | [pg_lexo_17-0.6.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lexo_17-0.6.1-1PGSTY.el8.x86_64.rpm) |
| `pg_lexo_17` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 265.6 KiB | [pg_lexo_17-0.6.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lexo_17-0.6.1-1PGSTY.el8.aarch64.rpm) |
| `pg_lexo_17` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 395.4 KiB | [pg_lexo_17-0.6.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lexo_17-0.6.1-1PGSTY.el9.x86_64.rpm) |
| `pg_lexo_17` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 285.1 KiB | [pg_lexo_17-0.6.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lexo_17-0.6.1-1PGSTY.el9.aarch64.rpm) |
| `pg_lexo_17` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 395.5 KiB | [pg_lexo_17-0.6.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lexo_17-0.6.1-1PGSTY.el10.x86_64.rpm) |
| `pg_lexo_17` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 285.1 KiB | [pg_lexo_17-0.6.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lexo_17-0.6.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-lexo` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 323.6 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 215.9 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 323.4 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 215.7 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 364.1 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 250.9 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 357.8 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 250.0 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 356.4 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-lexo` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 249.5 KiB | [postgresql-17-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-17-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lexo_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 378.4 KiB | [pg_lexo_16-0.6.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lexo_16-0.6.1-1PGSTY.el8.x86_64.rpm) |
| `pg_lexo_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 265.6 KiB | [pg_lexo_16-0.6.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lexo_16-0.6.1-1PGSTY.el8.aarch64.rpm) |
| `pg_lexo_16` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 395.4 KiB | [pg_lexo_16-0.6.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lexo_16-0.6.1-1PGSTY.el9.x86_64.rpm) |
| `pg_lexo_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 284.9 KiB | [pg_lexo_16-0.6.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lexo_16-0.6.1-1PGSTY.el9.aarch64.rpm) |
| `pg_lexo_16` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 395.4 KiB | [pg_lexo_16-0.6.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lexo_16-0.6.1-1PGSTY.el10.x86_64.rpm) |
| `pg_lexo_16` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 285.0 KiB | [pg_lexo_16-0.6.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lexo_16-0.6.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-lexo` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.5 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 222.6 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 329.9 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 222.5 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 370.8 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 257.5 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 364.5 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 256.6 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 362.9 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-lexo` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 256.0 KiB | [postgresql-16-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lexo/postgresql-16-pg-lexo_0.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Blad3Mak3r/pg_lexo" title="Repository" icon="github" subtitle="github.com/Blad3Mak3r/pg_lexo" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lexo-0.6.1.tar.gz pg_lexo-0.6.1-Cargo.lock" />}}
{{< /cards >}}


```bash
pig build pkg pg_lexo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_lexo;		# install via package name, for the active PG version

pig install pg_lexo -v 18;   # install for PG 18
pig install pg_lexo -v 17;   # install for PG 17
pig install pg_lexo -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lexo;
```

## Usage

Sources:

- [Official README](https://github.com/blad3mak3r/pg_lexo/blob/v0.6.1/README.md)
- [Version 0.6 SQL notes](https://github.com/blad3mak3r/pg_lexo/blob/v0.6.1/sql/pg_lexo--0.6.0.sql)
- [Public function implementation](https://github.com/blad3mak3r/pg_lexo/blob/v0.6.1/src/schema.rs)

`pg_lexo` 0.6.0 provides a Base62 `lexo` type for stable user-controlled ordering. It generates short values before, after, or between existing positions, allowing an item to move without renumbering every row. The type has comparison operators plus default B-tree and hash operator classes.

### Core Workflow

```sql
CREATE EXTENSION pg_lexo;

CREATE TABLE items (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL,
    position lexo NOT NULL UNIQUE DEFERRABLE
);

INSERT INTO items (title, position)
VALUES ('first', lexo_first());

INSERT INTO items (title, position)
SELECT 'last', lexo_after(position)
FROM items ORDER BY position DESC LIMIT 1;

INSERT INTO items (title, position)
VALUES ('middle', lexo_between('H'::lexo, 'I'::lexo));

SELECT * FROM items ORDER BY position;
```

`lexo_first()` starts at `H`. `lexo_after(lexo)`, `lexo_before(lexo)`, and `lexo_between(lexo, lexo)` generate neighboring positions; either argument to `lexo_between` may be NULL to represent an open end. Inputs contain only `0-9`, `A-Z`, and `a-z`, whose ASCII order matches the type's ordering.

### Table Helpers

`lexo_next(table, position_column, key_column, key_value)` finds the current maximum and returns the next position, optionally within a group. `lexo_add_column(table, column)` adds a `lexo` column. `lexo_rebalance(table, position_column, key_column, key_value)` rewrites matching rows with evenly distributed positions and returns the number updated.

The helpers issue dynamic SQL and require the caller to have the corresponding table privileges. Rebalancing updates every selected row and should run with application writers coordinated. If positions have a unique constraint, use a deferrable constraint and defer it within the rebalancing transaction to allow intermediate values to overlap. Position generation itself does not reserve a value: concurrent sessions can derive the same position, so enforce a suitable unique constraint and retry or serialize conflicting moves. When no Base62 value exists between adjacent positions (or before the minimum), rebalance first.

### Version 0.6 Boundary

Version 0.6 removed the dedicated `lexo` schema. The former `lexo.lexorank` type became `lexo`, and schema-qualified functions such as `lexo.first()` became names such as `lexo_first()`. Older examples are incompatible until migrated. Release/package 0.6.1 retains SQL/control version 0.6.0 and targets PostgreSQL 16–18 through pgrx 0.16.1. Pigsty packages patch `lexo_next` to use native ordering instead of the unavailable `max(lexo)` aggregate, and keep `lexo_rebalance` ordering independent of text collation.
