---
title: "pg_living_assertions"
linkTitle: "pg_living_assertions"
description: "Executable SQL checks with verdict dates and assertion history"
weight: 5300
categories: ["ADMIN"]
languages: ["SQL"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_living_assertions**](https://github.com/Manuelreyesbravo/pg_living_assertions) : Executable SQL checks with verdict dates and assertion history


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5300** | {{< badge content="pg_living_assertions" link="https://github.com/Manuelreyesbravo/pg_living_assertions" >}} | {{< ext "pg_living_assertions" >}} | `0.5.10` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `living_assertions` |
|    **Need By**    | {{< ext "pg_grammar_guard" >}} |

> [!Note] On-demand SQL checks run in a read-only subtransaction that is always rolled back; not per-write SQL ASSERTION constraints.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_living_assertions` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.10` | {{< bg "18" "pg_living_assertions_18" "green" >}} {{< bg "17" "pg_living_assertions_17" "green" >}} {{< bg "16" "pg_living_assertions_16" "green" >}} {{< bg "15" "pg_living_assertions_15" "green" >}} {{< bg "14" "pg_living_assertions_14" "green" >}} | `pg_living_assertions_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.10` | {{< bg "18" "postgresql-18-pg-living-assertions" "green" >}} {{< bg "17" "postgresql-17-pg-living-assertions" "green" >}} {{< bg "16" "postgresql-16-pg-living-assertions" "green" >}} {{< bg "15" "postgresql-15-pg-living-assertions" "green" >}} {{< bg "14" "postgresql-14-pg-living-assertions" "green" >}} | `postgresql-$v-pg-living-assertions` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Manuelreyesbravo/pg_living_assertions" title="Repository" icon="github" subtitle="github.com/Manuelreyesbravo/pg_living_assertions" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_living_assertions-0.5.10.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_living_assertions;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_living_assertions;		# install via package name, for the active PG version

pig install pg_living_assertions -v 18;   # install for PG 18
pig install pg_living_assertions -v 17;   # install for PG 17
pig install pg_living_assertions -v 16;   # install for PG 16
pig install pg_living_assertions -v 15;   # install for PG 15
pig install pg_living_assertions -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_living_assertions;
```

## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/README.md)
- [pg_living_assertions.control](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/CHANGELOG.md)
- [pg_living_assertions--0.5.9--0.5.10.sql](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions--0.5.9--0.5.10.sql)

`pg_living_assertions` 0.5.10 stores named SQL assertions, results and replacement history. Checks run on demand; they are not constraints evaluated on every write. This pure SQL extension needs no preload.

### Core Workflow

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### Operational Boundaries

A check returns one boolean `holds` and optional text `detail`. `living_assertions.run_all()` evaluates the registry; `living_assertions.state()`, `living_assertions.status` and `living_assertions.stale()` distinguish failures, unknown/erroring results and stale or absent checks. `living_assertions.declare_unchanged()` fingerprints a canonicalized expression; a value disappearing to NULL is broken. Checks run as their recorded author in a read-only, rolled-back subtransaction with the author search path and temporary schemas last. The caller must be allowed to `SET ROLE` to that author. Session advisory locks acquired during evaluation are released, but external effects and temporary-sequence changes are not a sandboxed resource. Only trusted authors should register SQL. Registration and checks are server-dated, tables reject TRUNCATE, and retirement/replacement is restricted to the author or an authorized role. Upgrade through the complete matching SQL chain; old assertions retain their existing fingerprint algorithm. Current releases use Apache-2.0; releases through 0.5.3 retain their original PostgreSQL License.

### Role Isolation in 0.5.10

Version 0.5.10 evaluates a check in a temporary `SECURITY DEFINER` function owned by its recorded author when role isolation is needed. PostgreSQL then refuses attempts to reset the role or session authorization inside that frame; the frame is created and rolled back with the read-only check. The author needs `TEMP` on the database when a frame is required, and the caller still needs authority to act as that author. A `SECURITY DEFINER` caller can now run another author's checks under that authority. This closes the role-reset escape in 0.5.8/0.5.9; it does not make external effects safe.
